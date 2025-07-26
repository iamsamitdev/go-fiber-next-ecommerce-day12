import Dashboard from '@/components/admin/dashboard'
import React from 'react'
import { Metadata } from "next"

export const metadata: Metadata = {
  title: "แดชบอร์ด | ร้านค้าออนไลน์",
  description: "ระบบจัดการร้านค้าออนไลน์สำหรับผู้ดูแลระบบ",
}

function DashboardPage() {
  return <Dashboard />
}

export default DashboardPage