(defun simpleoai-on-region (instruction)
  (call-process-region
   (min (mark) (point)) (max (mark) (point))
   "simpleoai"
   t
   t
   nil
   "-instruction" instruction
   ))

(defun simpleoai/do-instruction (instruction)
  (interactive "sInstruction: ")
  (simpleoai-on-region instruction))

(defun simpleoai/fix-english ()
  (interactive)
  (simpleoai-on-region "Keeping the input as close as possible to the output, including the style, fix grammar and orthographic errors"))

(defun simpleoai/reword-academically ()
  (interactive)
  (simpleoai-on-region "Keeping the input as close as possible to the output, reword it in such a way that it follows the conventions of academic English, but has the same meaning; keep it concise"))

(defun simpleoai/make-concise ()
  (interactive)
  (simpleoai-on-region "Keeping the style and meaning of the input, make it more concise"))
