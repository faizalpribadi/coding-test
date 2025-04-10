import { useState, useEffect } from "react";
import { useRouter } from "next/router";
import Head from 'next/head';
import Navigation from '../../components/Navigation';
import RepDetails from '../../components/RepDetails';
import Link from 'next/link';

export default function RepDetailsPage() {
  const router = useRouter();
  const { id } = router.query;
  
  const [rep, setRep] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    // Only fetch data when id is available (after hydration)
    if (!id) return;

    fetch("http://localhost:8000/api/sales-reps")
      .then((res) => res.json())
      .then((data) => {
        const foundRep = data.find(rep => rep.id.toString() === id);
        if (foundRep) {
          setRep(foundRep);
        } else {
          console.error("Rep not found");
        }
        setLoading(false);
      })
      .catch((err) => {
        console.error("Failed to fetch data:", err);
        setLoading(false);
      });
  }, [id]);

  // Format currency
  const formatCurrency = (value) => {
    return new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency: 'USD',
      maximumFractionDigits: 0
    }).format(value);
  };

  return (
    <div>
      <Head>
        <title>{rep ? `${rep.name} - Details` : 'Loading...'}</title>
        <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css" rel="stylesheet" />
      </Head>

      <Navigation />

      <main className="p-4">
        {loading ? (
          <div className="text-center p-4">
            <p>Loading rep details...</p>
          </div>
        ) : rep ? (
          <RepDetails
            selectedRep={rep}
            formatCurrency={formatCurrency}
          />
        ) : (
          <div className="text-center p-4">
            <p>Sales representative not found.</p>
            <Link href="/">
              <a className="border p-1 mt-2 inline-block">Back to Dashboard</a>
            </Link>
          </div>
        )}
      </main>
    </div>
  );
}
