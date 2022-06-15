import create from 'zustand'

export const useAuth = create((set) => ({
    isLogin: false,
    user: {},
    login: (user) => {
        set({
            isLogin: true,
            user: user
        })
    },
    doLogout: () => set({ isLogin: false }),
}))