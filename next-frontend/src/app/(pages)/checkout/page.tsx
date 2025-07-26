import React from "react"
import Checkout from "@/components/checkout"

import { Metadata } from "next"
export const metadata: Metadata = {
  title: "ตรวจสอบคำสั่งซื้อ | ร้านค้าออนไลน์",
  description: "ตรวจสอบคำสั่งซื้อ",
  // other metadata
}

const CheckoutPage = () => {
  return (
    <main>
      <Checkout />
    </main>
  )
}

export default CheckoutPage
