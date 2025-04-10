import React from 'react';

export default function AIAssistant({ 
  question, 
  setQuestion, 
  handleAskQuestion, 
  aiLoading, 
  answer 
}) {
  return (
    <div>
      <h1 className="mb-4">AI Assistant</h1>

      <div className="border p-3 mb-4">
        <h2 className="mb-2">Ask a Question</h2>
        <p className="mb-2">
          Ask me anything
        </p>

        <div className="flex flex-col md:flex-row gap-2">
          <input
            type="text"
            placeholder="Enter your question..."
            value={question}
            onChange={(e) => setQuestion(e.target.value)}
            onKeyPress={(e) => e.key === 'Enter' && handleAskQuestion()}
            className="border p-1 flex-grow"
          />
          <button
            onClick={handleAskQuestion}
            disabled={aiLoading}
            className="border p-1"
          >
            {aiLoading ? 'Thinking...' : 'Ask'}
          </button>
        </div>

        {answer && (
          <div className="border-t mt-3 pt-3">
            <h3 className="mb-1">AI Response:</h3>
            <div className="whitespace-pre-line">{answer}</div>
          </div>
        )}
      </div>
    </div>
  );
}
