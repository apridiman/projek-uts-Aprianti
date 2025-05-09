# Gunakan image Nginx sebagai base image
FROM nginx:alpine

# Hapus file default Nginx yang ada
RUN rm -rf /usr/share/nginx/html/*

# Salin file HTML, CSS, dan JS ke direktori Nginx
COPY . /usr/share/nginx/html

# Expose port 80 untuk akses HTTP
EXPOSE 80

# Jalankan Nginx di background
CMD ["nginx", "-g", "daemon off;"]
