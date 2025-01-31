import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import Link from "next/link";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
         <div className="border-r-2 w-[200px]  h-screen fixed  " >
          
          <Link href="/dashboard"> 
          <img className="p-4 bg-white" src="/Financfy-Logo-for-web.svg" alt=""  />
          </Link>
        
        
        <div  className=" bg-primary  p-6  pt-5 h-screen ">
        
          
    <ul >
      <li className="mb-2 text-lg text-white"><Link href="/dashboard">Dashboard</Link></li>
      <li className="mb-2 text-lg text-white">Transaction </li>
      <li className="mb-2 pl-7 pt-2 text-sm text-white font-medium"><Link href="/cashFlow">Cash Flow</Link></li>
      <li className="mb-2 pl-7 pt-2 text-sm text-white font-medium"><Link href="/cashIn">Cash In</Link></li>
      <li className="mb-2 pl-7 text-sm text-white font-medium"><Link href="/cashOut">Cash Out</Link></li>
      <li className="mb-2 pl-7 text-sm text-white font-medium"><Link href="/bulkEntries">Bulk Entries</Link></li>
    </ul>
  </div>
        </div>
        <header className="flex justify-between items-center ml-48 pl-12 pr-6 py-5 bg-white">
                      <div>
                         <h1 className="text-sm text-black font-bold ">Hello, Mst Rukaiya Islam Tonni! 👋</h1>
                         <p className="text-black text-xs">26 January 2025</p>
                     </div>
                <div className="flex justify-between gap-4">
                      <button className="px-4 py-2 bg-primary text-sm rounded-md">
                          + Add new cash book
                      </button>
                      <img className='bg-white h-10 w-10' src="/login_logo.png" alt="" />
          
                </div>
                
                </header>
        
        {children}
      </body>
    </html>
  );
}
