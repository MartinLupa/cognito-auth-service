"use client"

import { resendOTP } from "@/actions";
import { useTransition } from "react";

export function ResendOtpButton({ email }: { email: string }) {
 const [isPending, startTransition] = useTransition()

 const handleClick = () => {
  startTransition(async () => {
   await resendOTP(email)
  })
 }

 return (
  <button
   onClick={handleClick}
   disabled={isPending}
   className="text-sm text-gray-400 underline hover:text-gray-100 disabled:text-gray-400"
  >
   {isPending ? 'Resending...' : 'Resend OTP'}
  </button>
 );
}