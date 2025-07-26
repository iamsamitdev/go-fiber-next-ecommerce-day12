import React from "react"
import { Wishlist } from "@/components/wishlist"
import { Metadata } from "next"

export const metadata: Metadata = {
  title: "รายการสิ่งที่อยากได้ | ร้านค้าออนไลน์",
  description: "รายการสิ่งที่อยากได้",
  // other metadata
}

const WishlistPage = () => {
  return (
    <main>
      <Wishlist />
    </main>
  )
}

export default WishlistPage
