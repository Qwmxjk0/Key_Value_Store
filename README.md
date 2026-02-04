# go102: Append-only KV Store (ฝึก Go + พื้นฐาน DB storage)

โปรเจกต์นี้เป็น KV store แบบ append-only เพื่อฝึก Go และแนวคิดพื้นฐานของ storage engine
เน้นการเรียนรู้ด้วยการออกแบบและทดสอบทีละโดเมน

## เป้าหมายการเรียน
- เข้าใจ record encoding/decoding
- ทำ scan file เพื่อ rebuild in-memory index
- เขียน test 3 แบบ: write/read, restart/rebuild, corruption

## สเปกที่ตกลงแล้ว
- key เป็น binary, max 256 bytes
- value max 64KB
- record format:
  - `| key_len (u32) | val_len (u32) | key bytes | value bytes | checksum (u32) |`
- checksum = CRC32 (คำนวณจาก `key_len + val_len + key + value`)
- ทุกค่าตัวเลขหลาย byte ใช้ little-endian
- append-only file
- มี in-memory index (`key -> offset`)

## ขอบเขตตอนนี้
- เน้น single-file append-only
- โฟกัส correctness ก่อน performance
- ยังไม่ครอบคลุมเรื่อง compaction, transactions, หรือ replication

## แนวทางการพัฒนา (สรุป)
1. ออกแบบ API และ data flow
2. นิยาม record encoding/decoding
3. สแกนไฟล์เพื่อ rebuild index
4. เขียนและรันชุดทดสอบหลัก

## การทดสอบ (planned)
- write/read: เขียนแล้วอ่านตรงตาม key
- restart/rebuild: ปิดแล้วเปิดใหม่ สแกนไฟล์เพื่อสร้าง index
- corruption: ตรวจจับ checksum ผิด/ข้อมูลไม่ครบ

## พฤติกรรมหลักที่ตัดสินใจแล้ว
- `Open()` จะสร้างไฟล์ `data.db` หากยังไม่มี
- `Open()` จะ rebuild in-memory index จากไฟล์ทุกครั้ง
- หากพบ `ErrChecksum` / `ErrInvalidLength` / `ErrEOF` ระหว่าง rebuild → `Open()` จะคืน error

## เอกสารและการทำงานร่วมกัน
- หลีกเลี่ยงการเดา หากไม่ชัดเจนให้ระบุว่า UNKNOWN
- แยก FACT / ASSUMPTION / UNKNOWN ให้ชัดในการสื่อสาร
- ทุกการทดสอบควรสรุปผลและข้อจำกัดไว้เสมอ
