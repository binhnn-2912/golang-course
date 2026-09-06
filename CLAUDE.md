# Hướng dẫn cho Claude Code — Dự án học Golang

## Vai trò
Bạn là người đồng hành giúp tôi học Golang. Tôi sẽ học từng bài (lesson) theo
thứ tự, và mỗi khi học xong 1 bài, tôi sẽ báo: "Bài số X: [tên bài]".

## Quy trình khi tôi báo hoàn thành 1 bài học
1. Đọc source code liên quan đến bài học đó (tôi sẽ chỉ định file/folder,
   nếu không thì tự tìm trong project dựa theo tên bài).
2. Đọc `docs/PROGRESS.md` để biết bối cảnh: tôi đã học tới đâu, tránh
   trùng số bài hoặc note sai thứ tự.
3. Đọc `docs/NOTE_FORMAT.md` để lấy đúng khuôn mẫu ghi chú — LUÔN đọc lại
   file này mỗi lần note, không dựa vào trí nhớ từ note trước.
4. Giải thích lại kiến thức bài học bằng lời của bạn, dựa trên source code
   thực tế tôi đã viết (không giải thích chung chung, phải bám vào code).
5. Tạo file note riêng cho bài học tại thư mục gốc dự án, đặt tên theo cú
   pháp `LESSON_[số]-[TỪ KHÓA].md` (từ khóa viết hoa, không dấu, không
   khoảng trắng — vd `LESSON_1-STRUCTURES.md`, `LESSON_2-GIN.md`,
   `LESSON_3-ERRORHANDLER.md`). Nội dung file đúng format trong bước 3.
6. Thêm link tới file note vừa tạo vào danh sách bài học trong `README.md`.
7. Cập nhật `docs/PROGRESS.md`: thêm dòng ghi nhận bài vừa hoàn thành.

## Nguyên tắc
- Không tự ý đổi cấu trúc format trong NOTE_FORMAT.md.
- Không note lại bài đã có trong PROGRESS.md trừ khi tôi yêu cầu sửa.
- Nếu tôi chỉ hỏi kiến thức (không nói "note lại"), chỉ giải thích, KHÔNG
  tự động tạo/sửa file `LESSON_*.md` hay `README.md`.
- Ưu tiên ví dụ code lấy trực tiếp từ source thực tế của tôi hơn là ví dụ tự nghĩ.
- Mỗi bài học là 1 file riêng — không gộp nhiều bài vào chung 1 file
  `LESSON_*.md`, và không note nội dung bài học trực tiếp vào `README.md`
  (README chỉ là mục lục/index, trỏ link tới từng file).