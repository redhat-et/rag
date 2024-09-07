import os
import shutil
import chromadb
import logging
from pathlib import Path

from langchain_text_splitters import MarkdownHeaderTextSplitter

_log = logging.getLogger(__name__)

def gather_docs(path):
    docs = {}
    directory_path = Path(path)
    converted_doc_paths = list(directory_path.glob('*.md'))
    for doc in converted_doc_paths:
        with open(doc, 'r', encoding='utf-8') as file:
            markdown_content = file.read()
            docs[doc] = markdown_content
    return docs

def init_collection(): 
    try:
        client = chromadb.Client()
    except Exception as error:
        _log.error(f"Could not connect to the db: {error}")
    
    ten_k_collection = client.get_or_create_collection("10-Ks")
    return ten_k_collection

def chunk_document(document):
    headers_to_split_on = [
        ("#", "Header 1"),
        ("##", "Header 2"),
        ("###", "Header 3"),
    ]

    markdown_splitter = MarkdownHeaderTextSplitter(headers_to_split_on=headers_to_split_on)
    md_header_splits = markdown_splitter.split_text(document)
    return md_header_splits


def main():
    ten_k_collection = init_collection()
    docs = gather_docs("./pdf-output")
    for pdf_name, pdf_content in docs.items(): # iterate through every actual document
        doc_chunks = chunk_document(pdf_content)
        chunk_ids = []
        metadatas = []
        documents = []
        for chunk_index in range(len(doc_chunks)):
            chunk_ids.append(f"{pdf_name}.{chunk_index}")
            if not bool(doc_chunks[chunk_index].metadata):
                metadatas.append({"source": str(pdf_name)})
            else:
                doc_chunks[chunk_index].metadata["source"] = str(pdf_name)
                metadatas.append(doc_chunks[chunk_index].metadata)
            documents.append(doc_chunks[chunk_index].page_content)
        ten_k_collection.add(documents=documents, ids=chunk_ids, metadatas=metadatas)

if __name__ == "__main__":
    main()
