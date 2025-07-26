"use client"
import React, { useEffect } from "react"
import SingleItem from "./SingleItem"
import Image from "next/image"
import Link from "next/link"
import { useProduct } from "@/hooks/useProduct"

const BestSeller = () => {
  const { products, isLoading, error, fetchProducts } = useProduct()

  useEffect(() => {
    // Fetch 8 latest products with the lowest stock
    fetchProducts(1, 8, "stock", "asc")
  }, [fetchProducts])

  return (
    <section className="overflow-hidden">
      <div className="max-w-[1170px] w-full mx-auto px-4 sm:px-8 xl:px-0">
        {/* <!-- section title --> */}
        <div className="mb-10 flex items-center justify-between">
          <div>
            <span className="flex items-center gap-2.5 font-medium text-dark mb-1.5">
              <Image
                src="/images/icons/icon-07.svg"
                alt="icon"
                width={17}
                height={17}
              />
              เดือนนี้
            </span>
            <h2 className="font-semibold text-xl xl:text-heading-5 text-dark">
              สินค้าขายดี
            </h2>
          </div>
        </div>

        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-7.5">
          {isLoading ? (
            <p>Loading...</p>
          ) : error ? (
            <p>Error: {error}</p>
          ) : (
            products.map((item) => <SingleItem item={item} key={item.id} />)
          )}
        </div>

        <div className="text-center mt-12.5">
          <Link
            href="/shop-with-sidebar"
            className="inline-flex font-medium text-custom-sm py-3 px-7 sm:px-12.5 rounded-md border-gray-3 border bg-gray-1 text-dark ease-out duration-200 hover:bg-dark hover:text-white hover:border-transparent"
          >
            ดูทั้งหมด
          </Link>
        </div>
      </div>
    </section>
  )
}

export default BestSeller
