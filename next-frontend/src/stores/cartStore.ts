import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";
import { Product } from "@/types/product";

export interface CartItem extends Product {
  quantity: number;
}

interface CartState {
  cartItems: CartItem[];
  addToCart: (item: Product) => void;
  removeFromCart: (productId: string) => void;
  updateQuantity: (productId: string, quantity: number) => void;
  clearCart: () => void;
  cartTotal: () => number;
  cartItemCount: () => number;
}

const useCartStore = create<CartState>()(
  persist(
    (set, get) => ({
      cartItems: [],

      // ฟังก์ชันเพิ่มสินค้าไปยังตะกร้า
      addToCart: (item) => {
        set((state) => {
          const existingItem = state.cartItems.find(
            (cartItem: CartItem) => cartItem.id === item.id
          );
          // ถ้ามีสินค้าในตะกร้าแล้ว เพิ่มจำนวน
          if (existingItem) {
            return {
              cartItems: state.cartItems.map((cartItem) =>
                cartItem.id === item.id
                  ? { ...cartItem, quantity: cartItem.quantity + 1 }
                  : cartItem
              ),
            };
          }
          // ถ้ายังไม่มีสินค้าในตะกร้า เพิ่มสินค้าใหม่
          return { cartItems: [...state.cartItems, { ...item, quantity: 1 }] };
        });
      },

      // ฟังก์ชันลบสินค้าออกจากตะกร้า
      removeFromCart: (productId) =>
        set((state) => ({
          cartItems: state.cartItems.filter(
            (cartItem) => cartItem.id !== productId
          ),
        })),

      // ฟังก์ชันอัปเดตจำนวนสินค้าในตะกร้า
      updateQuantity: (productId, quantity) =>
        set((state) => ({
          cartItems: state.cartItems.map((cartItem) =>
            cartItem.id === productId ? { ...cartItem, quantity } : cartItem
          ),
        })),

      // ฟังก์ชันล้างตะกร้า
      clearCart: () => set({ cartItems: [] }),

      // ฟังก์ชันคำนวณราคารวมของตะกร้า
      cartTotal: () =>
        get().cartItems.reduce(
          (total, item) => total + item.price * item.quantity,
          0
        ),

      // ฟังก์ชันนับจำนวนสินค้าทั้งหมดในตะกร้า
      cartItemCount: () =>
        get().cartItems.reduce((total, item) => total + item.quantity, 0),
    }),
    {
      name: "cart-storage",
      storage: createJSONStorage(() => localStorage),
    }
  )
);

export default useCartStore;
