import { create } from "zustand"
import { persist, createJSONStorage } from "zustand/middleware"
import { Product } from "@/types/product"
import { toast } from "react-hot-toast"

interface WishlistItem extends Product {
  quantity: number
}

interface WishlistState {
  wishlist: WishlistItem[]
  addToWishlist: (item: Product) => void
  removeFromWishlist: (itemId: string) => void
  clearWishlist: () => void
  isInWishlist: (itemId: string) => boolean
}

const useWishlistStore = create<WishlistState>()(
  persist(
    (set, get) => ({
      wishlist: [],
      addToWishlist: (item: Product) => {
        const { wishlist } = get()
        const existingItem = wishlist.find((i) => i.id === item.id)

        if (existingItem) {
          toast("สินค้านี้อยู่ในรายการที่ชอบแล้ว")
          return
        }
        
        const newItem = { ...item, quantity: 1 }
        set({ wishlist: [...wishlist, newItem] })
        toast.success("เพิ่มในรายการที่ชอบแล้ว")
      },
      removeFromWishlist: (itemId: string) => {
        set((state) => ({
          wishlist: state.wishlist.filter((item) => item.id !== itemId),
        }))
        toast.success("ลบออกจากรายการที่ชอบแล้ว")
      },
      clearWishlist: () => {
        set({ wishlist: [] })
      },
      isInWishlist: (itemId: string) => {
        const { wishlist } = get()
        return wishlist.some((item) => item.id === itemId)
      }
    }),
    {
      name: "wishlist-storage",
      storage: createJSONStorage(() => localStorage),
    }
  )
)

export default useWishlistStore 