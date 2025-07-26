import React from "react"
import ShopWithSidebar from "@/components/shopwithsidebar"
import { Metadata } from "next"

export const metadata: Metadata = {
  title: "ร้านค้า | E-commerce Project",
  description: "This is Home FiberNext Project",
  // other metadata
}

const ShopWithSidebarPage = () => {
  return (
    <main>
      <ShopWithSidebar />
    </main>
  )
}

export default ShopWithSidebarPage
