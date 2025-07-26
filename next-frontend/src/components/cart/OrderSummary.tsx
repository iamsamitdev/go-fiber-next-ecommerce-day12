"use client"

import React, { useState, useEffect } from "react"
import useCartStore from "@/stores/cartStore"
import Link from "next/link"
import { formatPrice } from "@/lib/utils"

const OrderSummary = () => {
  const { cartItems, cartTotal } = useCartStore()
  const [isClient, setIsClient] = useState(false)

  useEffect(() => {
    setIsClient(true)
  }, [])

  if (!isClient) {
    return null // หรือจะแสดง Skeleton UI ก็ได้
  }

  return (
    <div className="lg:max-w-[455px] w-full">
      {/* <!-- order list box --> */}
      <div className="bg-white shadow-1 rounded-[10px]">
        <div className="border-b border-gray-3 py-5 px-4 sm:px-8.5">
          <h3 className="font-medium text-xl text-dark">สรุปราคาสินค้า</h3>
        </div>

        <div className="pt-2.5 pb-8.5 px-4 sm:px-8.5">
          {/* <!-- title --> */}
          <div className="flex items-center justify-between py-5 border-b border-gray-3">
            <div>
              <h4 className="font-medium text-dark">สินค้า</h4>
            </div>
            <div>
              <h4 className="font-medium text-dark text-right">ราคาสินค้า</h4>
            </div>
          </div>

          {/* <!-- product item --> */}
          {cartItems.map((item) => (
            <div
              key={item.id}
              className="flex items-center justify-between py-5 border-b border-gray-3"
            >
              <div>
                <p className="text-dark">
                  {item.title} x{item.quantity}
                </p>
              </div>
              <div>
                <p className="text-dark text-right">
                  {formatPrice(item.price * item.quantity)}
                </p>
              </div>
            </div>
          ))}

          {/* <!-- total --> */}
          <div className="flex items-center justify-between pt-5">
            <div>
              <p className="font-medium text-lg text-dark">ราคารวม</p>
            </div>
            <div>
              <p className="font-medium text-lg text-dark text-right">
                {formatPrice(cartTotal())}
              </p>
            </div>
          </div>

          {/* <!-- checkout button --> */}
          <Link
            href="/checkout"
            className="w-full flex justify-center font-medium text-white bg-blue py-3 px-6 rounded-md ease-out duration-200 hover:bg-blue-dark mt-7.5"
          >
            ดำเนินการชำระเงิน
          </Link>
        </div>
      </div>
    </div>
  )
}

export default OrderSummary
