import React from "react"
import Cart from "@/components/cart"

import { Metadata } from "next"
export const metadata: Metadata = {
  title: "ตะกร้าสินค้า | ร้านค้าออนไลน์",
  description: "ตะกร้าสินค้า",
  // other metadata
}

const CartPage = () => {
  return (
    <>
      <Cart />
    </>
  )
}

export default CartPage
