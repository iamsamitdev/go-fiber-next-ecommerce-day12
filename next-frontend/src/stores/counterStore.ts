import { create } from "zustand"
import { persist, createJSONStorage } from "zustand/middleware"

const useCounterStore = create(
  persist(
    (set) => ({
      // Initial state
      count: 0,
      // Actions to modify the state
      inc: () => set((state: { count: number }) => ({ count: state.count + 1 })),
      dec: () => set((state: { count: number }) => ({ count: state.count - 1 })),
    }),
    {
      name: "counter-storage", // Name of the item in the storage
      storage: createJSONStorage(() => localStorage), // Use localStorage as the storage
    }
  )
)

export default useCounterStore
