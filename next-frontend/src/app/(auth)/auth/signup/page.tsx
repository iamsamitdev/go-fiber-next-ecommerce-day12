import Signup from "@/components/auth/signup"
import React from "react"
import { Metadata } from "next"

export const metadata: Metadata = {
  title: "สมัครสมาชิก | ร้านค้าออนไลน์",
  description: "สมัครสมาชิก",
}

const SignupPage = () => {
  return <Signup />
}

export default SignupPage 