import React from 'react';

export default function Stats({ stats, formatCurrency }) {
  return (
    <div className="grid grid-cols-1 md:grid-cols-4 gap-2 mb-4">
      <div className="border p-3">
        <div>
          <h3>Total Sales Value</h3>
          <p>{formatCurrency(stats.totalValue)}</p>
        </div>
      </div>

      <div className="border p-3">
        <div>
          <h3>Closed Won</h3>
          <p>{formatCurrency(stats.closedWon)}</p>
        </div>
      </div>

      <div className="border p-3">
        <div>
          <h3>In Progress</h3>
          <p>{formatCurrency(stats.inProgress)}</p>
        </div>
      </div>

      <div className="border p-3">
        <div>
          <h3>Closed Lost</h3>
          <p>{formatCurrency(stats.closedLost)}</p>
        </div>
      </div>
    </div>
  );
}
