import create from "zustand";
import { persist } from "zustand/middleware";
import axios from "axios";
import { baseUrl } from "../Constant";

export const useCategory = create(
  persist(
    (set) => ({
      categories: [],
      fetch: async () => {
        try {
          const { data } = await axios.get(baseUrl + "/api/categories",{
            headers:{
                'Authorization': localStorage.getItem('token')
              }
          });
          set({
            categories: data,
          });
        } catch (error) {
          console.log(error);
        }
      },
    }),
    {
      name: "categories",
    }
  )
);
