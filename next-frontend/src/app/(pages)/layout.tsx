import type { Metadata } from "next"
import { Anuphan } from "next/font/google"
import "@/app/globals.css"
import Header from "@/components/shared/site/Header"
import Footer from "@/components/shared/site/Footer"
import { ModalProvider } from "@/context/QuickViewModalContext"
import { PreviewSliderProvider } from "@/context/PreviewSliderContext"
import QuickViewModal from "@/components/common/QuickViewModal"
import PreviewSliderModal from "@/components/common/PreviewSlider"
import { CartModalProvider } from "@/context/CartSidebarModalContext"
import CartSidebarModal from "@/components/common/CartSidebarModal"
  import { Toaster } from "react-hot-toast"

const anuphan = Anuphan({
  variable: "--font-anuphan",
  subsets: ["latin"],
})

export const metadata: Metadata = {
  title: "ร้านค้า | E-commerce Project",
  description: "This is Home FiberNext Project",
  // other metadata
}

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <html lang="en">
      <body
        className={`${anuphan.variable}`}
      >
        <CartModalProvider>
          <ModalProvider>
            <PreviewSliderProvider>
              <Header />
              {children}
              <QuickViewModal />
              <PreviewSliderModal />
              <CartSidebarModal />
              <Toaster />
            </PreviewSliderProvider>
          </ModalProvider>
        </CartModalProvider>
        <Footer />
      </body>
    </html>
  )
}
