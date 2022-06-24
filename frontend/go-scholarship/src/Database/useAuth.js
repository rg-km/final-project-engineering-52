import create from "zustand";
import { persist } from "zustand/middleware";
import axios from "axios";
import { baseUrl } from "../Constant";

export const useAuth = create(
  persist(
    (set) => ({
      user: {},
      login: async (user, navigate) => {
        try {
          const { data } = await axios.post(baseUrl + "/login", user);
          set({
            user: data.data,
          });
          localStorage.setItem("token", data.token);
          navigate("/");
        } catch (error) {
          console.log(error);
        }
      },
      doLogout: (navigate) => {
        localStorage.removeItem("token");
        set({ user: {} });
        navigate("/");
      },
    }),
    {
      name: "beasiswa-user",
    }
  )
);
