## 1. Thông tin
- Tool xóa firewall ảo từ vcenter, dựa vào tên firewall và trạng thái (On/Off)
- Viết  bằng Go Lang.
- Sử dụng ứng dụng có sẵn của go [govc](https://github.com/vmware/govmomi/tree/master/govc).
## 2. Hướng dẫn cài đặt
- [Cài đặt golang](https://go.dev/doc/install)
- Cài đăt govc.
```
go install github.com/vmware/govmomi/govc@latest
```
- Thêm vào các biến môi trường sau:
```
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$HOME/.local/bin:$PATH
export GO111MODULE=on
export GOVC_INSECURE=true
export GOVC_URL="https://vcenter.vietnix.vn"
export GOVC_USERNAME=<username vcenter>
export GOVC_PASSWORD=<password vcenter>
```

- Clone folder hiện tại về máy cá nhân, chạy chương trình
```
go run main.go
```

*Lưu ý: Chạy chương trình ở cty, nếu sử dụng ở nhà cần cài đặt tunnel (chưa fix được cách xử dụng proxychains để chạy)*

## 3. Ý tưởng viết code
- List được tất cả VM trong 1 datacenter, sau khi đã cài đặt đầy đủ môi trường:
```
govc ls <path>
```
![Hình minh họa](https://i.imgur.com/AJhn40m.png)
- Duyệt qua từng VM, tách lấy những VM có chứa cụm từ **-off-** và kết thúc bằng pattern **dd-mm-yyyy** hoặc **dd-mm-yy** để xử lý ở bước tiếp theo.
- Có đường dẫn tuyệt đối của một VM, chúng ta dễ dàng list được thông tin của VM đó, ở đây thấy được rất nhiều thông tin hữu ích, chúng ta chỉ tập trung vào trường **Power state**, xét Power state đang ở powerOff là sẽ thõa mãn điều kiện thứ 2, sẵn sàng để xóa.

![Hình minh họa 2](https://i.imgur.com/QNdverQ.png)

- Xét thỏa mãn 2 điều kiện trên ta sẽ đưa vào lệnh xóa.

## 4. Hạn chế.
- Hàm xóa của **govc** khá nguy hiểm vì chỉ cần truyền vào path của vm là nó sẽ **off con VM đó xóa nó** nên cần phải xét chính xác đường dẫn của VM và thõa mãn trạng thái **powerOff** thì mới đưa vào hàm xóa *(chưa tìm đc giải pháp để thay thế và test chi tiết - vì ở hạ tầng).*
- Còn phụ thuộc hoàn toàn vào ứng dụng viết sẵn bên thứ 3.
- Cơ chế ghi log còn hơi nghèo nàn (dùng tạm).
- Chủ yếu dựa vào bên thứ 3 nên hoàn toàn có thể viết bằng ngôn ngữ khác (miễn là dễ dàng xử lý được chuỗi). 
- Cấu trúc project chưa chuẩn (do người viết vừa viết code vừa học ngôn ngữ).
