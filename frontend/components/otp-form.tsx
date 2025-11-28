"use client"

import { verifyOTPAction } from "@/actions"
import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {
  Field,
  FieldDescription,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field"
import {
  InputOTP,
  InputOTPGroup,
  InputOTPSlot,
} from "@/components/ui/input-otp"
import { redirect, useSearchParams } from "next/navigation"
import { useActionState } from "react"
import { Alert, AlertDescription, AlertTitle } from "./ui/alert"
import { Terminal } from "lucide-react"
import { FaCheck } from "react-icons/fa"

export function OTPForm({ ...props }: React.ComponentProps<typeof Card>) {
  const [state, formAction, isPending] = useActionState(verifyOTPAction, null)
  const searchParams = useSearchParams()
  const email = searchParams.get('email')

  if (state?.success) {
    setTimeout(() => {
      redirect("/signin")
    }, 3000);
  }

  return (
    <Card {...props}>
      {state?.success ? (
        <div className="w-full max-w-xl text-center px-6">
          <FaCheck className="mx-auto mb-4 text-4xl text-green-500" />
          <h1 className="mb-6 text-2xl font-bold">Your email has been verified successfully!</h1>
          <p>We are redirecting you...</p>
        </div>
      ) :
        <div>
          <CardHeader>
            <CardTitle>Enter verification code</CardTitle>
            {state?.error && (<Alert variant="destructive">
              <Terminal />
              <AlertTitle>Login error</AlertTitle>
              <AlertDescription>
                {state.error}
              </AlertDescription>
            </Alert>)}
            <CardDescription>We sent a 6-digit code to your email.</CardDescription>
          </CardHeader>
          <CardContent>
            <form action={formAction}>
              {email && (
                <input type="hidden" name="email" value={email} />
              )}
              <FieldGroup>
                <Field>
                  <FieldLabel htmlFor="otp">Verification code</FieldLabel>
                  <InputOTP maxLength={6} id="otp" name="otp" required>
                    <InputOTPGroup className="gap-2.5 *:data-[slot=input-otp-slot]:rounded-md *:data-[slot=input-otp-slot]:border">
                      <InputOTPSlot index={0} />
                      <InputOTPSlot index={1} />
                      <InputOTPSlot index={2} />
                      <InputOTPSlot index={3} />
                      <InputOTPSlot index={4} />
                      <InputOTPSlot index={5} />
                    </InputOTPGroup>
                  </InputOTP>
                  <FieldDescription>
                    Enter the 6-digit code sent to your email.
                  </FieldDescription>
                </Field>
                <FieldGroup>
                  <Button type="submit" disabled={isPending}>
                    {isPending ? 'Logging in...' : 'Login'}
                  </Button>
                  <FieldDescription className="text-center">
                    Didn&apos;t receive the code? <a href="#">Resend</a>
                  </FieldDescription>
                </FieldGroup>
              </FieldGroup>
            </form>
          </CardContent>
        </div>
      }
    </Card>
  )
}
