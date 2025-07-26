import Contact from "@/components/contact"

import { Metadata } from "next"
export const metadata: Metadata = {
  title: "ติดต่อเรา | ร้านค้าออนไลน์",
  description: "ติดต่อเรา",
  // other metadata
}

const ContactPage = () => {
  return (
    <main>
      <Contact />
    </main>
  )
}

export default ContactPage
