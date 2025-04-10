import React from 'react';

export default function Filters({ 
  regionFilter, 
  setRegionFilter, 
  statusFilter, 
  setStatusFilter, 
  regions, 
  applyFilters, 
  resetFilters 
}) {
  return (
    <div className="border p-3 mb-4">
      <h3 className="mb-2">Filter Sales Data</h3>
      <div className="flex flex-col md:flex-row gap-2">
        <div>
          <label htmlFor="region">Region</label>
          <select
            id="region"
            value={regionFilter}
            onChange={(e) => setRegionFilter(e.target.value)}
            className="border p-1 w-full"
          >
            <option value="">All Regions</option>
            {regions.map((region) => (
              <option key={region} value={region}>{region}</option>
            ))}
          </select>
        </div>

        <div>
          <label htmlFor="status">Deal Status</label>
          <select
            id="status"
            value={statusFilter}
            onChange={(e) => setStatusFilter(e.target.value)}
            className="border p-1 w-full"
          >
            <option value="">All Statuses</option>
            <option value="Closed Won">Closed Won</option>
            <option value="In Progress">In Progress</option>
            <option value="Closed Lost">Closed Lost</option>
          </select>
        </div>

        <div className="flex gap-2 mt-2 md:mt-auto">
          <button
            onClick={applyFilters}
            className="border p-1"
          >
            Apply Filters
          </button>
          <button
            onClick={resetFilters}
            className="border p-1"
          >
            Reset
          </button>
        </div>
      </div>
    </div>
  );
}
