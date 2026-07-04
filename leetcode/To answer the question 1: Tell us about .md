To answer the question 1: Tell us about one AI/ML project you are most proud of.
I want to mention about the The Project: Agentic AI for Ad Generation at Emerge Tech

The Problem: The traditional advertising design process was a bottleneck. Designers were spending hours on manual work in software like Photoshop to create ad variations, which slowed down the production lifecycle significantly.

The Solution: I engineered an end-to-end "Agentic AI" system. Instead of simple text generation, I built a pipeline that transforms natural language prompts into fully rendered visual designs.

I used LangGraph to orchestrate complex workflows, ensuring the AI followed logical design principles and brand consistency. I leveraged LLMs (like Gemini/GPT) to generate code (HTML/TailwindCSS) rather than just static images, allowing for structured layouts.

I implemented a Headless Browser engine to render this code into high-fidelity visual assets automatically.

The Impact: This system drastically reduced the ad creation lifecycle from hours of manual work to just minutes. It enabled the team to iterate on designs rapidly by simply adjusting natural language prompts, effectively automating the path from concept to visual rendering.

Well for the question 2: Explain a complex AI/ML concept to a completely non-technical person.

I will pickup the work with RAG pipelines at NewTofu

The Concept I want to explain is: Retrieval-Augmented Generation

Imagine you are taking a very difficult exam. Standard AI (like ChatGPT) is like a student trying to answer questions purely from memory. They might be smart, but if they don't remember the specific fact, they might guess or 'hallucinate' an answer.

Retrieval-Augmented Generation is like letting that student take an open-book exam. When you ask a question, the AI first looks up the exact page in a textbook—or in your company's documents—to find the right information.

Then, it writes the answer based specifically on what it just read.

This ensures the answer isn't just creative, but factual and based on your actual data.

About the final question: Describe a situation where you had to take initiative without anyone telling you what to do.

For the final question:Describe a situation where you had to take initiative without anyone telling you what to do.

Situation: Fixing Low Accuracy in Food Classification at AMILI

What I did: While working on the gut microbiome analysis platform, I identified that our existing system for classifying food items into over 3,000 taxonomy classes was significantly underperforming, with an accuracy rate of only 52%. Recognizing that this low accuracy compromised the precision of our health insights, I took the initiative to completely re-engineer the classification approach rather than waiting for a directive to fix it. I designed and implemented a new system leveraging GPT-3 embeddings combined with fine-tuned BERT/T5 models.

The Result: My proactive overhaul resulted in a massive performance leap, improving the classification accuracy from 52% to 92%. This success allowed us to deliver a reliable end-to-end AI solution for precision health analysis.