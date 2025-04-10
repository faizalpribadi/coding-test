import React from 'react';
import Stats from './Stats';
import Filters from './Filters';
import SalesTable from './SalesTable';

export default function Dashboard({
  stats,
  formatCurrency,
  regionFilter,
  setRegionFilter,
  statusFilter,
  setStatusFilter,
  regions,
  applyFilters,
  resetFilters,
  loading,
  filteredReps,
  salesReps
}) {
  return (
    <div>
      <h1 className="mb-4">Sales Dashboard</h1>

      <Stats stats={stats} formatCurrency={formatCurrency} />

      <Filters
        regionFilter={regionFilter}
        setRegionFilter={setRegionFilter}
        statusFilter={statusFilter}
        setStatusFilter={setStatusFilter}
        regions={regions}
        applyFilters={applyFilters}
        resetFilters={resetFilters}
      />

      <SalesTable
        loading={loading}
        filteredReps={filteredReps}
        salesReps={salesReps}
        resetFilters={resetFilters}
        formatCurrency={formatCurrency}
      />
    </div>
  );
}
