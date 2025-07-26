import React from "react"
import BlogDetailsWithSidebar from "@/components/blogdetailswithsidebar"

import { Metadata } from "next"
export const metadata: Metadata = {
  title: "รายละเอียดบล็อก | ร้านค้าออนไลน์",
  description: "รายละเอียดบล็อก",
  // other metadata
}

const BlogDetailsWithSidebarPage = () => {
  return (
    <main>
      <BlogDetailsWithSidebar />
    </main>
  )
}

export default BlogDetailsWithSidebarPage
