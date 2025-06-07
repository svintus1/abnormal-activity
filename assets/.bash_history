ls /mnt/
ls /media/
ls /
lsblk 
mount sda1 /mnt/SP
mkdir /mnt/SP
mount sda1 /mnt/SP
ls /mnt/
ls /mnt/SP/
mount /dev/sda1 /mnt/SP
ls /mnt/SP/
ls /media/
ls
cat anaconda-ks.cfg 
ls /
ls /run/media/
ls /run/media/svintus/SP\ Stream\ S05/
cp -r /run/media/svintus/SP\ Stream\ S05/backup /home/svintus/
ls
sexit
exit
cd /home/svintus/Downloads/
LIBGL_ALWAYS_SOFTWARE=1 ./Hiddify-Linux-x64.AppImage
dnf install akmod-nvidia xorg-x11-drv-nvidia-cuda
GDK_BACKEND=x11 ./Hiddify-Linux-x64.AppImage
exit
