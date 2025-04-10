import React from 'react';
import Link from 'next/link';

export default function RepDetails({ selectedRep, formatCurrency }) {
  return (
    <div>
      <Link href="/">
        Back to Dashboard
      </Link>

      <div className="border p-3 mb-4">
        <h3>{selectedRep.name}</h3>
        <p>{selectedRep.role} • {selectedRep.region}</p>

        <div className="mt-2">
          <p><strong>Skills:</strong></p>
          <div>
            {selectedRep.skills.map((skill, index) => (
              <span key={index} className="mr-1">
                {skill}{index < selectedRep.skills.length - 1 ? ', ' : ''}
              </span>
            ))}
          </div>
        </div>
      </div>

      <div className="border p-3 mb-4">
        <h3 className="mb-2">Deals</h3>
        <table className="w-full border-collapse border">
          <thead>
            <tr>
              <th className="border p-1 text-left">Client</th>
              <th className="border p-1 text-left">Value</th>
              <th className="border p-1 text-left">Status</th>
            </tr>
          </thead>
          <tbody>
            {selectedRep.deals.map((deal, index) => (
              <tr key={index}>
                <td className="border p-1">{deal.client}</td>
                <td className="border p-1">{formatCurrency(deal.value)}</td>
                <td className="border p-1">{deal.status}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      <div className="border p-3">
        <h3 className="mb-2">Clients</h3>
        <table className="w-full border-collapse border">
          <thead>
            <tr>
              <th className="border p-1 text-left">Name</th>
              <th className="border p-1 text-left">Industry</th>
              <th className="border p-1 text-left">Contact</th>
            </tr>
          </thead>
          <tbody>
            {selectedRep.clients.map((client, index) => (
              <tr key={index}>
                <td className="border p-1">{client.name}</td>
                <td className="border p-1">{client.industry}</td>
                <td className="border p-1">
                  <a href={`mailto:${client.contact}`}>
                    {client.contact}
                  </a>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
