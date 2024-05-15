# Check if arguments are provided
if [ $# -ne 2 ]; then
    echo "Usage: $0 <url> <destination_directory>"
    exit 1
fi

url=$1
destination_dir=$2

# Create destination directory if it doesn't exist
mkdir -p "$destination_dir"

# Download tar.gz file
echo "Downloading $url..."
wget -qO- "$url" | tar -xz -C "$destination_dir"

echo "Extraction complete."