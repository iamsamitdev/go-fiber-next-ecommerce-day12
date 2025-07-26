import React from "react"
import BlogGridWithSidebar from "@/components/bloggrid/index"

import { Metadata } from "next"
export const metadata: Metadata = {
  title: "บล็อกกริด | ร้านค้าออนไลน์",
  description: "บล็อกกริด",
  // other metadata
}

const BlogGridWithSidebarPage = () => {
  return (
    <>
      <BlogGridWithSidebar />
    </>
  )
}

export default BlogGridWithSidebarPage
