"use client"

import React, { useEffect, useState } from 'react'

// Types สำหรับ Stats Data
interface SalesStats {
  total_sales: number
  total_orders: number
  today_sales: number
  today_orders: number
  monthly_sales: number
  monthly_orders: number
  yearly_sales: number
  yearly_orders: number
}

interface ProductStats {
  total_products: number
  low_stock_products: number
  out_of_stock_products: number
  total_categories: number
}

interface UserStats {
  total_users: number
  active_users: number
  inactive_users: number
  new_users: number
}

interface DashboardData {
  salesStats: SalesStats | null
  productStats: ProductStats | null
  userStats: UserStats | null
  recentOrders: any[]
  loading: boolean
  error: string | null
}

const Dashboard = () => {
  const [data, setData] = useState<DashboardData>({
    salesStats: null,
    productStats: null,
    userStats: null,
    recentOrders: [],
    loading: true,
    error: null
  })

  useEffect(() => {
    fetchDashboardData()
  }, [])

  const fetchDashboardData = async () => {
    try {
      setData(prev => ({ ...prev, loading: true, error: null }))

      // จำลองการดึงข้อมูลจาก API (แทนที่ด้วย API จริงภายหลัง)
      const mockSalesStats: SalesStats = {
        total_sales: 150000,
        total_orders: 156,
        today_sales: 5000,
        today_orders: 12,
        monthly_sales: 45000,
        monthly_orders: 89,
        yearly_sales: 120000,
        yearly_orders: 145
      }

      const mockProductStats: ProductStats = {
        total_products: 248,
        low_stock_products: 12,
        out_of_stock_products: 3,
        total_categories: 15
      }

      const mockUserStats: UserStats = {
        total_users: 1250,
        active_users: 850,
        inactive_users: 400,
        new_users: 45
      }

      const mockRecentOrders = [
        { id: 1, customer: "สมชาย ใจดี", total: 2500, status: "pending", date: "2024-01-15" },
        { id: 2, customer: "สมหญิง สวยงาม", total: 1800, status: "completed", date: "2024-01-15" },
        { id: 3, customer: "วิชัย มั่นใจ", total: 3200, status: "processing", date: "2024-01-14" },
        { id: 4, customer: "นารี น่ารัก", total: 1500, status: "completed", date: "2024-01-14" },
        { id: 5, customer: "ชาตรี ชัยชนะ", total: 4100, status: "pending", date: "2024-01-13" }
      ]

      // จำลองการ delay API
      await new Promise(resolve => setTimeout(resolve, 1000))

      setData({
        salesStats: mockSalesStats,
        productStats: mockProductStats,
        userStats: mockUserStats,
        recentOrders: mockRecentOrders,
        loading: false,
        error: null
      })
    } catch (error) {
      setData(prev => ({
        ...prev,
        loading: false,
        error: 'เกิดข้อผิดพลาดในการดึงข้อมูล'
      }))
    }
  }

  const formatCurrency = (amount: number) => {
    return new Intl.NumberFormat('th-TH', {
      style: 'currency',
      currency: 'THB'
    }).format(amount)
  }

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'completed': return 'bg-green-100 text-green-700'
      case 'pending': return 'bg-yellow-100 text-yellow-700'
      case 'processing': return 'bg-blue-100 text-blue-700'
      case 'cancelled': return 'bg-red-100 text-red-700'
      default: return 'bg-gray-100 text-gray-700'
    }
  }

  const getStatusText = (status: string) => {
    switch (status) {
      case 'completed': return 'เสร็จสิ้น'
      case 'pending': return 'รอดำเนินการ'
      case 'processing': return 'กำลังประมวลผล'
      case 'cancelled': return 'ยกเลิก'
      default: return status
    }
  }

  if (data.loading) {
    return (
      <div className="p-6">
        <div className="flex items-center justify-center h-96">
          <div className="animate-spin rounded-full h-16 w-16 border-b-2 border-blue-600"></div>
        </div>
      </div>
    )
  }

  if (data.error) {
    return (
      <div className="p-6">
        <div className="bg-red-50 border border-red-200 rounded-lg p-6 text-center">
          <h2 className="text-red-700 text-xl font-medium mb-2">เกิดข้อผิดพลาด</h2>
          <p className="text-red-600">{data.error}</p>
          <button 
            onClick={fetchDashboardData}
            className="mt-4 bg-red-600 text-white px-6 py-2 rounded-lg hover:bg-red-700 transition-colors"
          >
            ลองใหม่อีกครั้ง
          </button>
        </div>
      </div>
    )
  }

  return (
    <div className="p-6">
      <div className="max-w-7xl mx-auto">
        {/* Header */}
        <div className="flex flex-col sm:flex-row sm:items-center justify-between gap-4 mb-8">
          <div>
            <h1 className="font-semibold text-gray-900 text-2xl mb-2">
              แดชบอร์ดผู้ดูแลระบบ
            </h1>
            <p className="text-gray-600">ภาพรวมข้อมูลและสถิติการดำเนินงาน</p>
          </div>
          <div className="flex items-center gap-3">
            <button 
              onClick={fetchDashboardData}
              className="flex items-center gap-2 bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors"
            >
              <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
              </svg>
              รีเฟรช
            </button>
          </div>
        </div>

        {/* Stats Cards Grid */}
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
          {/* Total Sales Card */}
          <div className="bg-white rounded-xl shadow-sm border p-6">
            <div className="flex items-center justify-between mb-4">
              <div className="w-12 h-12 bg-blue-50 rounded-lg flex items-center justify-center">
                <svg className="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1" />
                </svg>
              </div>
              <span className="text-green-600 text-sm font-medium">+15.3%</span>
            </div>
            <h3 className="text-gray-600 text-sm mb-1">ยอดขายรวม</h3>
            <p className="text-2xl font-bold text-gray-900">{formatCurrency(data.salesStats?.total_sales || 0)}</p>
            <p className="text-xs text-gray-500 mt-1">คำสั่งซื้อทั้งหมด: {data.salesStats?.total_orders}</p>
          </div>

          {/* Today Sales Card */}
          <div className="bg-white rounded-xl shadow-sm border p-6">
            <div className="flex items-center justify-between mb-4">
              <div className="w-12 h-12 bg-green-50 rounded-lg flex items-center justify-center">
                <svg className="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
                </svg>
              </div>
              <span className="text-green-600 text-sm font-medium">+8.2%</span>
            </div>
            <h3 className="text-gray-600 text-sm mb-1">ยอดขายวันนี้</h3>
            <p className="text-2xl font-bold text-gray-900">{formatCurrency(data.salesStats?.today_sales || 0)}</p>
            <p className="text-xs text-gray-500 mt-1">คำสั่งซื้อวันนี้: {data.salesStats?.today_orders}</p>
          </div>

          {/* Total Products Card */}
          <div className="bg-white rounded-xl shadow-sm border p-6">
            <div className="flex items-center justify-between mb-4">
              <div className="w-12 h-12 bg-purple-50 rounded-lg flex items-center justify-center">
                <svg className="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                </svg>
              </div>
              <span className="text-red-600 text-sm font-medium">{data.productStats?.low_stock_products} ต่ำ</span>
            </div>
            <h3 className="text-gray-600 text-sm mb-1">สินค้าทั้งหมด</h3>
            <p className="text-2xl font-bold text-gray-900">{data.productStats?.total_products}</p>
            <p className="text-xs text-gray-500 mt-1">หมวดหมู่: {data.productStats?.total_categories}</p>
          </div>

          {/* Total Users Card */}
          <div className="bg-white rounded-xl shadow-sm border p-6">
            <div className="flex items-center justify-between mb-4">
              <div className="w-12 h-12 bg-orange-50 rounded-lg flex items-center justify-center">
                <svg className="w-6 h-6 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-.5a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1z" />
                </svg>
              </div>
              <span className="text-green-600 text-sm font-medium">+{data.userStats?.new_users} ใหม่</span>
            </div>
            <h3 className="text-gray-600 text-sm mb-1">ผู้ใช้ทั้งหมด</h3>
            <p className="text-2xl font-bold text-gray-900">{data.userStats?.total_users}</p>
            <p className="text-xs text-gray-500 mt-1">ใช้งาน: {data.userStats?.active_users}</p>
          </div>
        </div>

        {/* Content Grid */}
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          {/* Recent Orders */}
          <div className="lg:col-span-2 bg-white rounded-xl shadow-sm border">
            <div className="px-6 py-4 border-b border-gray-200">
              <div className="flex items-center justify-between">
                <h2 className="font-medium text-lg text-gray-900">คำสั่งซื้อล่าสุด</h2>
                <button className="text-blue-600 hover:text-blue-700 transition-colors text-sm">
                  ดูทั้งหมด
                </button>
              </div>
            </div>
            <div className="p-6">
              <div className="overflow-x-auto">
                <table className="w-full min-w-[500px]">
                  <thead>
                    <tr className="border-b border-gray-200">
                      <th className="text-left py-3 text-sm font-medium text-gray-600">ลูกค้า</th>
                      <th className="text-left py-3 text-sm font-medium text-gray-600">ยอดรวม</th>
                      <th className="text-left py-3 text-sm font-medium text-gray-600">สถานะ</th>
                      <th className="text-left py-3 text-sm font-medium text-gray-600">วันที่</th>
                    </tr>
                  </thead>
                  <tbody>
                    {data.recentOrders.map((order) => (
                      <tr key={order.id} className="border-b border-gray-100 hover:bg-gray-50 transition-colors">
                        <td className="py-4">
                          <div className="flex items-center gap-3">
                            <div className="w-8 h-8 bg-blue-50 rounded-full flex items-center justify-center">
                              <span className="text-blue-600 text-sm font-medium">
                                {order.customer.charAt(0)}
                              </span>
                            </div>
                            <span className="text-gray-900 font-medium">{order.customer}</span>
                          </div>
                        </td>
                        <td className="py-4 text-gray-900 font-medium">{formatCurrency(order.total)}</td>
                        <td className="py-4">
                          <span className={`inline-flex px-2 py-1 rounded-full text-xs font-medium ${getStatusColor(order.status)}`}>
                            {getStatusText(order.status)}
                          </span>
                        </td>
                        <td className="py-4 text-gray-600">{new Date(order.date).toLocaleDateString('th-TH')}</td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            </div>
          </div>

          {/* Quick Stats & Actions */}
          <div className="space-y-6">
            {/* Monthly Performance */}
            <div className="bg-white rounded-xl shadow-sm border p-6">
              <h3 className="font-medium text-gray-900 mb-4">ประสิทธิภาพรายเดือน</h3>
              <div className="space-y-4">
                <div className="flex items-center justify-between">
                  <span className="text-gray-600">ยอดขาย</span>
                  <span className="font-medium text-gray-900">{formatCurrency(data.salesStats?.monthly_sales || 0)}</span>
                </div>
                <div className="flex items-center justify-between">
                  <span className="text-gray-600">คำสั่งซื้อ</span>
                  <span className="font-medium text-gray-900">{data.salesStats?.monthly_orders}</span>
                </div>
                <div className="flex items-center justify-between">
                  <span className="text-gray-600">สินค้าคงเหลือน้อย</span>
                  <span className="font-medium text-red-600">{data.productStats?.low_stock_products}</span>
                </div>
                <div className="flex items-center justify-between">
                  <span className="text-gray-600">สินค้าหมด</span>
                  <span className="font-medium text-red-600">{data.productStats?.out_of_stock_products}</span>
                </div>
              </div>
            </div>

            {/* Quick Actions */}
            <div className="bg-white rounded-xl shadow-sm border p-6">
              <h3 className="font-medium text-gray-900 mb-4">การจัดการด่วน</h3>
              <div className="space-y-3">
                <button className="w-full flex items-center gap-3 p-3 text-left rounded-lg hover:bg-gray-50 transition-colors">
                  <div className="w-8 h-8 bg-blue-50 rounded-lg flex items-center justify-center">
                    <svg className="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                    </svg>
                  </div>
                  <div>
                    <p className="font-medium text-gray-900">เพิ่มสินค้าใหม่</p>
                    <p className="text-xs text-gray-600">จัดการสินค้าในระบบ</p>
                  </div>
                </button>

                <button className="w-full flex items-center gap-3 p-3 text-left rounded-lg hover:bg-gray-50 transition-colors">
                  <div className="w-8 h-8 bg-green-50 rounded-lg flex items-center justify-center">
                    <svg className="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                    </svg>
                  </div>
                  <div>
                    <p className="font-medium text-gray-900">จัดการคำสั่งซื้อ</p>
                    <p className="text-xs text-gray-600">ตรวจสอบและอัปเดตสถานะ</p>
                  </div>
                </button>

                <button className="w-full flex items-center gap-3 p-3 text-left rounded-lg hover:bg-gray-50 transition-colors">
                  <div className="w-8 h-8 bg-purple-50 rounded-lg flex items-center justify-center">
                    <svg className="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                    </svg>
                  </div>
                  <div>
                    <p className="font-medium text-gray-900">จัดการผู้ใช้</p>
                    <p className="text-xs text-gray-600">ดูและแก้ไขข้อมูลผู้ใช้</p>
                  </div>
                </button>

                <button className="w-full flex items-center gap-3 p-3 text-left rounded-lg hover:bg-gray-50 transition-colors">
                  <div className="w-8 h-8 bg-orange-50 rounded-lg flex items-center justify-center">
                    <svg className="w-4 h-4 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v4a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                    </svg>
                  </div>
                  <div>
                    <p className="font-medium text-gray-900">รายงานและสถิติ</p>
                    <p className="text-xs text-gray-600">ดูข้อมูลเชิงลึกต่างๆ</p>
                  </div>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Dashboard
