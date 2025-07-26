"use client"

import React, { useEffect, useState } from "react"
import Breadcrumb from "@/components/common/Breadcrumb"
import SingleItem from "@/components/wishlist/SingleItem"
import useWishlistStore from "@/stores/wishlistStore"
import Link from "next/link"

export const Wishlist = () => {

  // Get wishlist items from the store
  const { wishlist, clearWishlist } = useWishlistStore()
  const [isClient, setIsClient] = useState(false)

  useEffect(() => {
    setIsClient(true)
  }, [])

  if (!isClient) {
    return null // ป้องกันการเรนเดอร์ฝั่งเซิร์ฟเวอร์
  }

  return (
    <>
      <Breadcrumb title={"สิ่งที่อยากได้"} pages={["สิ่งที่อยากได้"]} />
      <section className="overflow-hidden py-20 bg-gray-2">
        <div className="max-w-[1170px] w-full mx-auto px-4 sm:px-8 xl:px-0">
          <div className="flex flex-wrap items-center justify-between gap-5 mb-7.5">
            <h2 className="font-medium text-dark text-2xl">สิ่งที่อยากได้</h2>
            <button className="text-blue cursor-pointer" onClick={clearWishlist}>ล้างรายการสิ่งที่อยากได้</button>
          </div>

          <div className="bg-white rounded-[10px] shadow-1">
            <div className="w-full overflow-x-auto">
              <div className="min-w-[1170px]">
                {/* <!-- table header --> */}
                <div className="flex items-center py-5.5 px-10">
                  <div className="min-w-[83px]"></div>
                  <div className="min-w-[387px]">
                    <p className="text-dark">สินค้า</p>
                  </div>

                  <div className="min-w-[205px]">
                    <p className="text-dark">ราคา</p>
                  </div>

                  <div className="min-w-[265px]">
                    <p className="text-dark">สถานะ</p>
                  </div>

                  <div className="min-w-[150px]">
                    <p className="text-dark text-right">การดำเนินการ</p>
                  </div>
                </div>

                {/* <!-- wish item --> */}
                {wishlist.length > 0 ? (
                  wishlist.map((item, key) => (
                    <SingleItem
                      item={{
                        ...item,
                        inStock: (item.stock || 0) > 0,
                        imgs: {
                          thumbnails: item.imgs?.thumbnails || [],
                          previews: item.imgs?.previews || [],
                        },
                      }}
                      key={key}
                    />
                  ))
                ) : (
                  <div className="text-center py-20">
                    <h3 className="mb-4">รายการสิ่งที่อยากได้ของคุณว่างเปล่า</h3>
                    <Link
                      href="/shop-with-sidebar"
                      className="inline-flex font-medium text-custom-sm py-2.5 px-7 rounded-md border-gray-3 border bg-gray-1 text-dark ease-out duration-200 hover:bg-dark hover:text-white hover:border-transparent"
                    >
                      เลือกซื้อสินค้า
                    </Link>
                  </div>
                )}
              </div>
            </div>
          </div>
        </div>
      </section>
    </>
  )
}
