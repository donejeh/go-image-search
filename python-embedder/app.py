from flask import Flask, request, jsonify
from PIL import Image
import torch
from torchvision import transforms
import clip
import os

app = Flask(__name__)
device = "cuda" if torch.cuda.is_available() else "cpu"

model, preprocess = clip.load("ViT-B/32", device=device)

# Function for embedding images
@app.route("/embed", methods=["POST"])
def embed():
    data = request.json
    image_path = data.get("image_path")

    if not image_path or not os.path.exists(image_path):
        return jsonify({"error": "Invalid image path"}), 400

    image = preprocess(Image.open(image_path)).unsqueeze(0).to(device)
    with torch.no_grad():
        embedding = model.encode_image(image).cpu().numpy().tolist()[0]

    return jsonify({"embedding": embedding})

# Function for embedding text
def Text(text: str):
    text_tokens = clip.tokenize(text).to(device)
    with torch.no_grad():
        text_embedding = model.encode_text(text_tokens)
    return text_embedding


# Function for comparing image and text similarity
def Compare(image, text: str):
    print(text)
    image = preprocess(image).unsqueeze(0).to(device)
    text = clip.tokenize(text).to(device)
    
    with torch.no_grad():
        logits_per_image, logits_per_text = model(image, text)
        probs = logits_per_image.softmax(dim=-1).cpu().numpy()
        return np.ravel(probs)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
