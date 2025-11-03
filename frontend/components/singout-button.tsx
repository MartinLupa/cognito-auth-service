"use client"

import { signoutAction } from '@/actions';
import { useActionState } from 'react';
import { Button } from './ui/button';
import { redirect } from 'next/navigation';

export default function SignoutButton() {
 const [state, formAction, isPending] = useActionState(signoutAction, null)

 if (state?.success) {
  redirect('/signin')
 }

 return (
  <form action={formAction}>
   <Button type="submit" disabled={isPending}>Logout</Button>
  </form>
 )
}