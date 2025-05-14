import { Geist, Geist_Mono } from "next/font/google";
import { WebSocketProvider } from "@/context/wsContext";
import { UserProvider } from "@/context/userContext";

import "./globals.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata = {
  title: "zillcode",
  description:
    "zillcode is a platform for learning and sharing knowledge about programming, technology, and software development.",
  keywords:
    "zillcode, programming, technology, software development, learning, knowledge sharing",
  authors: [{ name: "zillcode", url: "https://zillcode.com" }],
  creators: [{ name: "zillcode", url: "https://zillcode.com" }],
  publisher: "zillcode",
  applicationName: "zillcode",
  icons: {
    icon: "/icon.svg",
    shortcut: "/icon.svg",
    apple: "/icon.svg",
  },
};

export const viewport = {
  themeColor: "#ffffff",
  colorScheme: "light dark",
  width: "device-width",
  initialScale: 1,
};

export default function RootLayout({ children }) {
  return (
    <UserProvider>
      <WebSocketProvider>
        <html lang="en">
          <body className={`${geistSans.variable} ${geistMono.variable}`}>
            {children}
          </body>
        </html>
      </WebSocketProvider>
    </UserProvider>
  );
}
