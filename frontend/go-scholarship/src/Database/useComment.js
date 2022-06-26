import create from "zustand";
import { persist } from "zustand/middleware";
import axios from "axios";
import { baseUrl } from "../Constant";

export const useComment = create(
  persist(
    (set) => ({
        comments: [],
      fetch: async () => {
        try {
          const { data } = await axios.get(baseUrl + "/api/comments",{
            headers:{
                'Authorization': localStorage.getItem('token')
              }
          });
          set({
            comments: data,
          });
        } catch (error) {
          console.log(error);
        }
      },
    }),
    {
      name: "comments",
    }
  )
);
