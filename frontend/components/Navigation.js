import React from 'react';
import Link from 'next/link';
import { useRouter } from 'next/router';

export default function Navigation() {
  const router = useRouter();

  return (
    <nav className="border-b p-3">
      <div className="flex">
        <div className="mr-4">
          <span className="font-bold">InterOpera</span>
        </div>
        <div className="flex gap-4">
          <Link href="/">
            Dashboard
          </Link>
          <Link href="/ai-assistant">
            AI Assistant
          </Link>
        </div>
      </div>
    </nav>
  );
}
