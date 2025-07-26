import { Category } from "@/hooks/useCategory"
import React from "react"
import Image from "next/image"

const SingleItem = ({ item }: { item: Category }) => {
  return (
    <a href="#" className="group flex flex-col items-center">
      <div className="max-w-[130px] w-full bg-[#F2F3F8] h-32.5 rounded-full flex items-center justify-center mb-4 overflow-hidden">
        <Image
          src={item.image}
          alt={item.name}
          width={130}
          height={130}
          className="w-full h-full object-cover transition-transform duration-300 group-hover:scale-110"
        />
      </div>

      <div className="flex justify-center">
        <h3 className="inline-block font-medium text-center text-dark bg-gradient-to-r from-blue to-blue bg-[length:0px_1px] bg-left-bottom bg-no-repeat transition-[background-size] duration-500 hover:bg-[length:100%_3px] group-hover:bg-[length:100%_1px] group-hover:text-blue">
          {item.name}
        </h3>
      </div>
    </a>
  )
}

export default SingleItem
