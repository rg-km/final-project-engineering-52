import create from "zustand";
import { persist } from "zustand/middleware";
import axios from "axios";
import { baseUrl } from "../Constant";

export const useUsers = create(
  persist(
    (set) => ({
      users: [],
      fetch: async () => {
        try {
          const { data } = await axios.get(baseUrl + "/api/users", {
            headers:{
              'Authorization': localStorage.getItem('token')
            }
          });
          set({
            users: data.users,
          });
        } catch (error) {
          console.log(error);
        }
      },
    }),
    {
      name: "users",
    }
  )
);
