import React from 'react';
import Link from 'next/link';

export default function SalesTable({
  loading,
  filteredReps,
  salesReps,
  resetFilters,
  formatCurrency
}) {
  return (
    <div className="border p-3">
      <h3>Sales Representatives</h3>
      <p>Showing {filteredReps.length} of {salesReps.length} representatives</p>

      {loading ? (
        <div className="text-center p-4">
          <p>Loading sales data...</p>
        </div>
      ) : filteredReps.length === 0 ? (
        <div className="text-center p-4">
          <p>No sales representatives match your filters.</p>
          <button
            onClick={resetFilters}
            className="border p-1 mt-2"
          >
            Clear Filters
          </button>
        </div>
      ) : (
        <div className="overflow-x-auto">
          <table className="w-full border-collapse border mt-2">
            <thead>
              <tr>
                <th className="border p-1 text-left">Representative</th>
                <th className="border p-1 text-left">Region</th>
                <th className="border p-1 text-left">Deals</th>
                <th className="border p-1 text-left">Total Value</th>
                <th className="border p-1 text-left">Actions</th>
              </tr>
            </thead>
            <tbody>
              {filteredReps?.map((rep) => {
                const repTotalValue = rep.deals.reduce((sum, deal) => sum + deal.value, 0);

                return (
                  <tr key={rep.id}>
                    <td className="border p-1">
                      <div>{rep.name}</div>
                      <div>{rep.role}</div>
                    </td>
                    <td className="border p-1">
                      {rep.region}
                    </td>
                    <td className="border p-1">
                      <div>
                        {rep.deals.map((deal, index) => (
                          <span key={index} className="mr-1">
                            {deal.status}{index < rep.deals.length - 1 ? ', ' : ''}
                          </span>
                        ))}
                      </div>
                    </td>
                    <td className="border p-1">
                      {formatCurrency(repTotalValue)}
                    </td>
                    <td className="border p-1">
                      <Link href={`/rep/${rep.id}`}>
                        View Details
                      </Link>
                    </td>
                  </tr>
                );
              })}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}
