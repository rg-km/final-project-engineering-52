import create from "zustand";
import { persist } from "zustand/middleware";
import axios from "axios";
import { baseUrl } from "../Constant";

export const useScholarship = create(
  persist(
    (set) => ({
      scolarship: [],
      fetch: async () => {
        try {
          const { data } = await axios.get(baseUrl + "/api/scholarships");
          set({
            scolarship: data.data,
          });
        } catch (error) {
          console.log(error);
        }
      },
    }),
    {
      name: "scholarship",
    }
  )
);
