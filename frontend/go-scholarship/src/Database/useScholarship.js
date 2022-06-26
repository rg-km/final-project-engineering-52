import create from "zustand";
import { persist } from "zustand/middleware";
import axios from "axios";
import { baseUrl } from "../Constant";
import Scholarship from "../Pages/Admin/Scholarship";

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
      delete_data: async (id) => {
        try {
           await axios.delete(baseUrl + "/api/scholarships/" + id);
          const data = Scholarship.filter(item => item.id != id)
          set({
            scolarship: data
          })
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
