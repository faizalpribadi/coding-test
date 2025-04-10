import { useState, useEffect } from "react";
import Head from 'next/head';
import Navigation from '../components/Navigation';
import Dashboard from '../components/Dashboard';
import RepDetails from '../components/RepDetails';

export default function Home() {
  const [salesReps, setSalesReps] = useState([]);
  const [filteredReps, setFilteredReps] = useState([]);
  const [loading, setLoading] = useState(true);
  const [stats, setStats] = useState({
    totalValue: 0,
    closedWon: 0,
    inProgress: 0,
    closedLost: 0
  });
  const [regionFilter, setRegionFilter] = useState('');
  const [statusFilter, setStatusFilter] = useState('');
  const [regions, setRegions] = useState([]);


  // Fetch sales data
  useEffect(() => {
    fetch("http://localhost:8000/api/sales-reps")
      .then((res) => res.json())
      .then((data) => {
        setSalesReps(data || []);
        setFilteredReps(data || []);
        setLoading(false);

        const uniqueRegions = [...new Set(data.map(rep => rep.region))];
        setRegions(uniqueRegions);

        calculateStats(data);
      })
      .catch((err) => {
        console.error("Failed to fetch data:", err);
        setLoading(false);
      });
  }, []);

  const calculateStats = (reps) => {
    let totalValue = 0;
    let closedWon = 0;
    let inProgress = 0;
    let closedLost = 0;

    reps.forEach(rep => {
      rep.deals.forEach(deal => {
        totalValue += deal.value;

        if (deal.status === "Closed Won") {
          closedWon += deal.value;
        } else if (deal.status === "In Progress") {
          inProgress += deal.value;
        } else if (deal.status === "Closed Lost") {
          closedLost += deal.value;
        }
      });
    });

    setStats({
      totalValue,
      closedWon,
      inProgress,
      closedLost
    });
  };

  const applyFilters = () => {
    let filtered = [...salesReps];

    if (regionFilter) {
      filtered = filtered.filter(rep => rep.region === regionFilter);
    }

    if (statusFilter) {
      filtered = filtered.filter(rep => {
        return rep.deals.some(deal => deal.status === statusFilter);
      });
    }

    setFilteredReps(filtered);
    calculateStats(filtered);
  };

  const resetFilters = () => {
    setRegionFilter('');
    setStatusFilter('');
    setFilteredReps(salesReps);
    calculateStats(salesReps);
  };

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
        <title>Sales Dashboard</title>
        <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css" rel="stylesheet" />
      </Head>

      <Navigation />

      <main className="p-4">
        <Dashboard
          stats={stats}
          formatCurrency={formatCurrency}
          regionFilter={regionFilter}
          setRegionFilter={setRegionFilter}
          statusFilter={statusFilter}
          setStatusFilter={setStatusFilter}
          regions={regions}
          applyFilters={applyFilters}
          resetFilters={resetFilters}
          loading={loading}
          filteredReps={filteredReps}
          salesReps={salesReps}
        />
      </main>
    </div>
  );
}
