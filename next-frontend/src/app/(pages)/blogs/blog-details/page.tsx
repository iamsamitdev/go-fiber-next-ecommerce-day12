import BlogDetails from "@/components/blogdetails"
import React from "react"

import { Metadata } from "next"
export const metadata: Metadata = {
  title: "รายละเอียดบล็อก | ร้านค้าออนไลน์",
  description: "รายละเอียดบล็อก",
  // other metadata
}

const BlogDetailsPage = () => {
  return (
    <main>
      <BlogDetails />
    </main>
  )
}

export default BlogDetailsPage
