import type { Metadata } from "next"
import { Anuphan } from "next/font/google"
import "@/app/globals.css"
import Link from "next/link"
import React from "react"

const anuphan = Anuphan({
  variable: "--font-anuphan",
  subsets: ["latin"],
})

export const metadata: Metadata = {
  title: "เข้าสู่ระบบ | ร้านค้าออนไลน์",
  description: "ระบบจัดการผู้ใช้",
}

interface AuthLayoutProps {
  children: React.ReactNode
}

export default function AuthLayout({ children }: AuthLayoutProps) {
  return (
    <html lang="th">
      <body className={`${anuphan.variable}`}>
        {/* Back to Home Button */}
        <Link 
          href="/" 
          className="fixed top-4 left-4 z-50 bg-white hover:bg-gray-100 rounded-full p-3 shadow-lg transition-all duration-200 hover:scale-110"
          title="กลับหน้าหลัก"
        >
          <svg 
            width="24" 
            height="24" 
            viewBox="0 0 24 24" 
            fill="none" 
            xmlns="http://www.w3.org/2000/svg"
            className="text-gray-700"
          >
            <path 
              d="M3 12L5 10M5 10L12 3L19 10M5 10V20C5 20.5523 5.44772 21 6 21H9M19 10L21 12M19 10V20C19 20.5523 18.5523 21 18 21H15M9 21C9.55228 21 10 20.5523 10 20V16C10 15.4477 10.4477 15 11 15H13C13.5523 15 14 15.4477 14 16V20C14 20.5523 14.4477 21 15 21M9 21H15" 
              stroke="currentColor" 
              strokeWidth="2" 
              strokeLinecap="round" 
              strokeLinejoin="round"
            />
          </svg>
        </Link>
        
        <div className="min-h-screen flex">
          {/* Left Column - Image & Text */}
          <div className="hidden lg:flex lg:w-1/2 bg-gradient-to-br from-blue-600 to-purple-700 relative">
            <div className="absolute inset-0">
              <img
                src="https://images.unsplash.com/photo-1556742049-0cfed4f6a45d?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=2340&q=80"
                alt="Shopping Experience"
                className="w-full h-full object-cover opacity-20"
              />
            </div>
            <div className="relative z-10 flex flex-col justify-center items-start text-left p-12 text-white">
              <div className="max-w-md">
                <h1 className="text-4xl font-bold mb-6">
                  ยินดีต้อนรับสู่ร้านค้าออนไลน์
                </h1>
                <p className="text-xl mb-8 leading-relaxed">
                  ค้นพบสินค้าคุณภาพดีที่คัดสรรมาเพื่อคุณ พร้อมประสบการณ์การช้อปปิ้งที่ไม่เหมือนใคร
                </p>
                <div className="space-y-4 text-left">
                  <div className="flex items-center space-x-3">
                    <div className="w-2 h-2 bg-white rounded-full"></div>
                    <span className="text-lg">สินค้าคุณภาพสูง</span>
                  </div>
                  <div className="flex items-center space-x-3">
                    <div className="w-2 h-2 bg-white rounded-full"></div>
                    <span className="text-lg">ส่งฟรีทั่วประเทศ</span>
                  </div>
                  <div className="flex items-center space-x-3">
                    <div className="w-2 h-2 bg-white rounded-full"></div>
                    <span className="text-lg">บริการหลังการขาย 24/7</span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          {/* Right Column - Form */}
          <div className="w-full lg:w-1/2 flex flex-col justify-center px-8 bg-gray-50 min-h-screen">
            <div className="max-w-md mx-auto w-full">              
              {children}
            </div>
          </div>
        </div>
      </body>
    </html>
  )
} 