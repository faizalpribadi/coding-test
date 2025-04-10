import { useState } from "react";
import Head from 'next/head';
import Navigation from '../components/Navigation';
import AIAssistant from '../components/AIAssistant';

export default function AIAssistantPage() {
  // State for AI
  const [question, setQuestion] = useState("");
  const [answer, setAnswer] = useState("");
  const [aiLoading, setAiLoading] = useState(false);

  // Handle AI questions
  const handleAskQuestion = async () => {
    if (!question.trim()) return;

    setAiLoading(true);
    setAnswer("");

    try {
      const response = await fetch("http://localhost:8000/api/ai", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ question }),
      });
      const data = await response.json();
      setAnswer(data.answer);
    } catch (error) {
      console.error("Error in AI request:", error);
      setAnswer("Sorry, there was an error processing your request.");
    } finally {
      setAiLoading(false);
    }
  };

  return (
    <div>
      <Head>
        <title>AI Assistant</title>
        <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css" rel="stylesheet" />
      </Head>

      <Navigation />

      <main className="p-4">
        <AIAssistant
          question={question}
          setQuestion={setQuestion}
          handleAskQuestion={handleAskQuestion}
          aiLoading={aiLoading}
          answer={answer}
        />
      </main>
    </div>
  );
}
