import Signin from "@/components/auth/signin"
import React from "react"
import { Metadata } from "next"

export const metadata: Metadata = {
  title: "เข้าสู่ระบบ | ร้านค้าออนไลน์",
  description: "เข้าสู่ระบบ",
}

const SigninPage = () => {
  return <Signin />
}

export default SigninPage 