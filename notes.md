Consideration: should we be using multi-modal?
    - Via docling we could export the tables themselves as images. Would this be a better format? It seems like a better storage format but not for queries. What is the best way for the model to get that kind of table data out?
        - https://github.com/GoogleCloudPlatform/generative-ai/blob/main/gemini/use-cases/retrieval-augmented-generation/multimodal_rag_langchain.ipynb has some examples
        - for our implementation this would leverage the docling export figures example: https://github.com/DS4SD/docling/blob/main/examples/export_figures.py
            - we would have to get the table images in the right inersertion order for the chunks
