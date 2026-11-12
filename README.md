# Go File Processing & Word Counter

Dự án thực hành xử lý file dung lượng lớn trong Go sử dụng `bufio.Scanner`, giúp đọc dữ liệu theo luồng và thống kê tần suất từ vựng mà không làm tràn bộ nhớ (RAM).

## Tính năng chính

- **Đọc file theo từng dòng (Line-by-Line Reader):** Duyệt tuần tự từng dòng dữ liệu thông qua vùng nhớ đệm cố định, tối ưu cho các file log hoặc văn bản lớn.
- **Đếm tần suất từ (Word Counter):** Tự động tách từng từ theo khoảng trắng, chuyển về chữ thường, loại bỏ dấu câu ở hai đầu và thống kê số lần xuất hiện.

## Cấu trúc thư mục

```text
.
├── main.go          # Điểm khởi chạy chương trình, gọi và kết hợp các hàm xử lý
├── reader.go        # Hàm đọc nội dung file theo từng dòng
├── counter.go       # Hàm tách từ, làm sạch dấu câu và đếm tần suất
├── test.txt         # File văn bản mẫu để chạy thử nghiệm
└── README.md        # Tài liệu hướng dẫn dự án
```

*(Lưu ý: Bạn có thể đổi tên `reader.go`, `counter.go` hoặc `test.txt` trong mục này cho khớp với tên file thực tế trên máy của bạn).*

## Yêu cầu hệ thống

- [Go](https://go.dev/dl/) phiên bản **1.16** trở lên.

## Cách chạy chương trình

1. **Clone repository về máy:**
   ```bash
   git clone https://github.com/<username>/<ten-repo>.git
   cd <ten-repo>
   ```

2. **Chạy chương trình:**
   Sử dụng lệnh sau để Go biên dịch cùng lúc `main.go` và các file chứa hàm trong thư mục:
   ```bash
   go run .
   ```
   *(Nếu chưa khởi tạo file `go.mod`, hãy dùng lệnh `go run *.go`)*.

## Ví dụ kết quả

Giả sử file `test.txt` có nội dung: `"Học Go rất thú vị. Go giúp xử lý file nhanh!"`, kết quả in ra màn hình sẽ là:

```text
học: 1 lần
go: 2 lần
rất: 1 lần
thú: 1 lần
vị: 1 lần
giúp: 1 lần
xử: 1 lần
lý: 1 lần
file: 1 lần
nhanh: 1 lần
```