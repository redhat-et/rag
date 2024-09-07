dbpath=$(realpath chromadb)
cd $dbpath
bash -c "chroma run --path $dbpath"
cd "$dbpath/.."
venv/bin/python 5_generate_embeddings_and_populate_vector_db.py
