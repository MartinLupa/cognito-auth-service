import { NextRequest, NextResponse } from "next/server";

const protectedPaths = ['/protected']

export async function middleware(request: NextRequest) {
 if (protectedPaths.some((path) => request.nextUrl.pathname.startsWith(path))) {
  const authCookie = request.cookies.get('session_token')

  if (!authCookie) {
   const loginUrl = new URL('/signin', request.url)
   return Response.redirect(loginUrl.toString())
  }

  try {
   const response = await fetch(new URL(process.env.AUTH_SERVICE_VALIDATE_SESSION_ENDPOINT || ''), {
    method: 'POST',
    headers: {
     'Content-Type': 'application/json',
     'Authorization': `${authCookie.value}`
    }
   })

   if (!response.ok) {
    const loginUrl = new URL('/signin', request.url)
    return Response.redirect(loginUrl.toString())
   }

  } catch (error) {
   console.error('Error validating JWT:', error)
   const loginUrl = new URL('/signin', request.url)
   return Response.redirect(loginUrl.toString())
  }

  return NextResponse.next()
 }
}