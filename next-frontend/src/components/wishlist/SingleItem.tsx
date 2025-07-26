import React from "react"
import Image from "next/image"
import { Product } from "@/types/product"
import useCartStore from "@/stores/cartStore"
import useWishlistStore from "@/stores/wishlistStore"
import { useCartModalContext } from "@/context/CartSidebarModalContext"
import { formatPrice } from "@/lib/utils"
import Link from "next/link"

interface WishlistItemFromProp extends Product {
  inStock: boolean
}

interface SingleItemProps {
  item: WishlistItemFromProp
}

const SingleItem = ({ item }: SingleItemProps) => {
  const { removeFromWishlist } = useWishlistStore()
  const { addToCart } = useCartStore()
  const { openCartModal } = useCartModalContext()

  const handleAddToCart = () => {
    addToCart(item)
    openCartModal()
  }

  return (
    <div className="flex items-center border-t border-gray-3 py-5 px-10">
      <div className="min-w-[83px]">
        <button
          onClick={() => removeFromWishlist(String(item.id))}
          aria-label="button for remove product from wishlist"
          className="flex items-center justify-center rounded-lg max-w-[38px] w-full h-9.5 bg-gray-2 border border-gray-3 ease-out duration-200 hover:bg-red-light-6 hover:border-red-light-4 hover:text-red"
        >
          <svg
            className="fill-current"
            width="22"
            height="22"
            viewBox="0 0 22 22"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M9.19509 8.22222C8.92661 7.95374 8.49131 7.95374 8.22282 8.22222C7.95433 8.49071 7.95433 8.92601 8.22282 9.1945L10.0284 11L8.22284 12.8056C7.95435 13.074 7.95435 13.5093 8.22284 13.7778C8.49133 14.0463 8.92663 14.0463 9.19511 13.7778L11.0006 11.9723L12.8061 13.7778C13.0746 14.0463 13.5099 14.0463 13.7784 13.7778C14.0469 13.5093 14.0469 13.074 13.7784 12.8055L11.9729 11L13.7784 9.19451C14.0469 8.92603 14.0469 8.49073 13.7784 8.22224C13.5099 7.95376 13.0746 7.95376 12.8062 8.22224L11.0006 10.0278L9.19509 8.22222Z"
              fill=""
            />
            <path
              fillRule="evenodd"
              clipRule="evenodd"
              d="M11.0007 1.14587C5.55835 1.14587 1.14648 5.55773 1.14648 11C1.14648 16.4423 5.55835 20.8542 11.0007 20.8542C16.443 20.8542 20.8548 16.4423 20.8548 11C20.8548 5.55773 16.443 1.14587 11.0007 1.14587ZM2.52148 11C2.52148 6.31713 6.31774 2.52087 11.0007 2.52087C15.6836 2.52087 19.4798 6.31713 19.4798 11C19.4798 15.683 15.6836 19.4792 11.0007 19.4792C6.31774 19.4792 2.52148 15.683 2.52148 11Z"
              fill=""
            />
          </svg>
        </button>
      </div>

      <div className="min-w-[387px]">
        <div className="flex items-center justify-between gap-5">
          <div className="w-full flex items-center gap-5.5">
            <div className="flex items-center justify-center rounded-[5px] bg-gray-2 max-w-[80px] w-full h-17.5 relative overflow-hidden">
              <Image
                src={item.imgs?.thumbnails[0] || "/images/products/product-1-bg-1.png"}
                alt={item.title}
                fill
                className="object-cover"
              />
            </div>

            <div>
              <Link href={`/shop-details/${item.id}`}>
                <h3 className="text-dark ease-out duration-200 hover:text-blue">
                  {item.title}
                </h3>
              </Link>
            </div>
          </div>
        </div>
      </div>

      <div className="min-w-[205px]">
        <p className="text-dark">{formatPrice(item.price)}</p>
      </div>

      <div className="min-w-[265px]">
        {item.inStock ? (
          <div className="flex items-center gap-1.5">
            <svg
              width="20"
              height="20"
              viewBox="0 0 20 20"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M10 1.25C5.16751 1.25 1.25 5.16751 1.25 10C1.25 14.8325 5.16751 18.75 10 18.75C14.8325 18.75 18.75 14.8325 18.75 10C18.75 5.16751 14.8325 1.25 10 1.25Z"
                fill="#10B981"
              />
              <path
                d="M13.7071 7.79289C14.0976 8.18342 14.0976 8.81658 13.7071 9.20711L9.20711 13.7071C8.81658 14.0976 8.18342 14.0976 7.79289 13.7071L6.29289 12.2071C5.90237 11.8166 5.90237 11.1834 6.29289 10.7929C6.68342 10.4024 7.31658 10.4024 7.70711 10.7929L8.5 11.5858L12.2929 7.79289C12.6834 7.40237 13.3166 7.40237 13.7071 7.79289Z"
                fill="white"
              />
            </svg>
            <span className="text-green-600">มีสินค้า</span>
          </div>
        ) : (
          <div className="flex items-center gap-1.5">
            <svg
              width="20"
              height="20"
              viewBox="0 0 20 20"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M9.99935 14.7917C10.3445 14.7917 10.6243 14.5119 10.6243 14.1667V9.16669C10.6243 8.82151 10.3445 8.54169 9.99935 8.54169C9.65417 8.54169 9.37435 8.82151 9.37435 9.16669V14.1667C9.37435 14.5119 9.65417 14.7917 9.99935 14.7917Z"
                fill="#F23030"
              />
              <path
                d="M9.99935 5.83335C10.4596 5.83335 10.8327 6.20645 10.8327 6.66669C10.8327 7.12692 10.4596 7.50002 9.99935 7.50002C9.53911 7.50002 9.16602 7.12692 9.16602 6.66669C9.16602 6.20645 9.53911 5.83335 9.99935 5.83335Z"
                fill="#F23030"
              />
              <path
                fillRule="evenodd"
                clipRule="evenodd"
                d="M1.04102 10C1.04102 5.05247 5.0518 1.04169 9.99935 1.04169C14.9469 1.04169 18.9577 5.05247 18.9577 10C18.9577 14.9476 14.9469 18.9584 9.99935 18.9584C5.0518 18.9584 1.04102 14.9476 1.04102 10ZM9.99935 2.29169C5.74215 2.29169 2.29102 5.74283 2.29102 10C2.29102 14.2572 5.74215 17.7084 9.99935 17.7084C14.2565 17.7084 17.7077 14.2572 17.7077 10C17.7077 5.74283 14.2565 2.29169 9.99935 2.29169Z"
                fill="#F23030"
              />
            </svg>
            <span className="text-red">หมดสต็อก</span>
          </div>
        )}
      </div>

      <div className="min-w-[150px] flex justify-end">
        <button
          onClick={handleAddToCart}
          disabled={!item.inStock}
          className={`inline-flex py-2.5 px-6 rounded-md ease-out duration-200 ${
            item.inStock
              ? "text-dark hover:text-white bg-gray-1 border border-gray-3 hover:bg-blue hover:border-gray-3"
              : "text-gray-400 bg-gray-100 border border-gray-200 cursor-not-allowed"
          }`}
        >
          เพิ่มลงตะกร้า
        </button>
      </div>
    </div>
  )
}

export default SingleItem
