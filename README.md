# Go Word Counter CLI

Công cụ dòng lệnh (CLI) viết bằng Go giúp đọc file văn bản dung lượng lớn theo từng từ (word-by-word streaming) để tiết kiệm RAM, thống kê tần suất xuất hiện và hiển thị danh sách các từ phổ biến nhất.

## Tính năng chính

* **Đọc file theo từng từ (Word-by-Word Streaming):** Sử dụng `bufio.Scanner` kết hợp `bufio.ScanWords` để duyệt tuần tự từng từ qua vùng nhớ đệm cố định, giúp xử lý file lớn mà không gây tràn bộ nhớ.
* **Đếm và xếp hạng Top từ (`CountWords` & `PrintTopWords`):** Chuẩn hóa chữ thường, loại bỏ dấu câu, thống kê tần suất và in ra nhóm các từ xuất hiện nhiều nhất.
* **Hỗ trợ tham số dòng lệnh (CLI Flags):** Cho phép truyền đường dẫn file và tùy chỉnh số lượng từ muốn hiển thị trực tiếp qua terminal, kèm kiểm tra tính hợp lệ của dữ liệu đầu vào.

## Cấu trúc thư mục

```text
.
├── main.go          # Xử lý CLI flags (-file, -top), mở file và điều phối chương trình
├── counter.go       # Chứa hàm CountWords: đọc từng từ và thống kê tần suất
├── printer.go       # Chứa hàm PrintTopWords: sắp xếp và hiển thị top N từ
├── test.txt         # File văn bản mẫu để chạy thử nghiệm
└── README.md        # Tài liệu hướng dẫn dự án
```

*(Lưu ý: Bạn có thể đổi tên `counter.go`, `printer.go` hoặc `test.txt` cho khớp với tên file thực tế trên máy của bạn).*

## Yêu cầu hệ thống

* [Go](https://go.dev/dl/) phiên bản **1.16** trở lên.

## Hướng dẫn sử dụng & Tham số (Flags)

Chương trình hỗ trợ các tham số dòng lệnh sau:

| Flag | Kiểu dữ liệu | Mặc định | Bắt buộc | Mô tả |
| :--- | :--- | :--- | :--- | :--- |
| `-file` | `string` | `""` | **Có** | Đường dẫn tới file văn bản cần thống kê. |
| `-top` | `int` | `5` | Không | Số lượng từ xuất hiện nhiều nhất muốn hiển thị (phải là số nguyên dương `> 0`). |

### Các lệnh chạy phổ biến

1. **Chạy với cấu hình mặc định (hiển thị top 5 từ):**
   ```bash
   go run . -file=test.txt
   ```

2. **Tùy chỉnh số lượng từ top muốn hiển thị (ví dụ top 10):**
   ```bash
   go run . -file=test.txt -top=10
   ```

3. **Xem hướng dẫn sử dụng các flag trực tiếp trên terminal:**
   ```bash
   go run . -h
   ```

*(Lưu ý: Nếu thư mục chưa khởi tạo file `go.mod`, hãy thay `go run .` bằng `go run *.go`)*.

## Ví dụ kết quả

Khi chạy lệnh `go run . -file=test.txt -top=3`:

```text
File path: test.txt
Top words to display: 3
go: 5
file: 3
scanner: 2
```