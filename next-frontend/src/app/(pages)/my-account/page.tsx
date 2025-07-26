import MyAccount from "@/components/myaccount"
import React from "react"

import { Metadata } from "next"
export const metadata: Metadata = {
  title: "บัญชีของฉัน | ร้านค้าออนไลน์",
  description: "บัญชีของฉัน",
  // other metadata
}

const MyAccountPage = () => {
  return (
    <main>
      <MyAccount />
    </main>
  )
}

export default MyAccountPage
