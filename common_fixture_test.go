package oak_test

import "github.com/phogolabs/parcello"

func addResourceWithMissingMigrationDir(manager *parcello.Manager) {
	manager.Add([]byte{
		80, 75, 3, 4, 20, 0, 8, 0, 8, 0, 53, 80, 154, 76, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 13, 0, 9, 0, 115, 99, 114, 105,
		112, 116, 47, 116, 101, 120, 116, 46, 116, 85, 84, 5, 0, 1,
		134, 163, 225, 90, 74, 68, 0, 46, 64, 0, 0, 0, 255, 255, 80,
		75, 7, 8, 166, 178, 131, 44, 10, 0, 0, 0, 12, 0, 0, 0, 80,
		75, 1, 2, 20, 3, 20, 0, 8, 0, 8, 0, 53, 80, 154, 76, 166,
		178, 131, 44, 10, 0, 0, 0, 12, 0, 0, 0, 13, 0, 9, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 164, 129, 0, 0, 0, 0, 115, 99, 114, 105,
		112, 116, 47, 116, 101, 120, 116, 46, 116, 85, 84, 5, 0, 1,
		134, 163, 225, 90, 80, 75, 5, 6, 0, 0, 0, 0, 1, 0, 1, 0, 68,
		0, 0, 0, 78, 0, 0, 0, 0, 0,
	})
}

func addResourceWithMissingMigrations(manager *parcello.Manager) {
	manager.Add([]byte{
		80, 75, 3, 4, 20, 0, 8, 0, 8, 0, 172, 80, 154, 76, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 34, 0, 9, 0, 109, 105, 103, 114,
		97, 116, 105, 111, 110, 47, 48, 48, 48, 54, 48, 53, 50, 52,
		48, 48, 48, 48, 48, 48, 95, 115, 101, 116, 117, 112, 46, 115,
		113, 108, 85, 84, 5, 0, 1, 100, 164, 225, 90, 1, 0, 0, 255,
		255, 80, 75, 7, 8, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 80,
		75, 3, 4, 20, 0, 8, 0, 8, 0, 174, 80, 154, 76, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 34, 0, 9, 0, 109, 105, 103, 114, 97,
		116, 105, 111, 110, 47, 50, 48, 49, 56, 48, 52, 48, 54, 49,
		57, 48, 48, 49, 53, 95, 117, 115, 101, 114, 115, 46, 115,
		113, 108, 85, 84, 5, 0, 1, 104, 164, 225, 90, 1, 0, 0, 255,
		255, 80, 75, 7, 8, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 80,
		75, 3, 4, 20, 0, 8, 0, 8, 0, 228, 80, 154, 76, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 13, 0, 9, 0, 115, 99, 114, 105, 112,
		116, 47, 116, 101, 115, 116, 46, 116, 85, 84, 5, 0, 1, 205,
		164, 225, 90, 1, 0, 0, 255, 255, 80, 75, 7, 8, 0, 0, 0, 0,
		5, 0, 0, 0, 0, 0, 0, 0, 80, 75, 1, 2, 20, 3, 20, 0, 8, 0,
		8, 0, 172, 80, 154, 76, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0,
		34, 0, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 164, 129, 0, 0, 0, 0,
		109, 105, 103, 114, 97, 116, 105, 111, 110, 47, 48, 48, 48,
		54, 48, 53, 50, 52, 48, 48, 48, 48, 48, 48, 95, 115, 101,
		116, 117, 112, 46, 115, 113, 108, 85, 84, 5, 0, 1, 100, 164,
		225, 90, 80, 75, 1, 2, 20, 3, 20, 0, 8, 0, 8, 0, 174, 80,
		154, 76, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 34, 0, 9, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 164, 129, 94, 0, 0, 0, 109, 105, 103,
		114, 97, 116, 105, 111, 110, 47, 50, 48, 49, 56, 48, 52, 48,
		54, 49, 57, 48, 48, 49, 53, 95, 117, 115, 101, 114, 115, 46,
		115, 113, 108, 85, 84, 5, 0, 1, 104, 164, 225, 90, 80, 75,
		1, 2, 20, 3, 20, 0, 8, 0, 8, 0, 228, 80, 154, 76, 0, 0, 0,
		0, 5, 0, 0, 0, 0, 0, 0, 0, 13, 0, 9, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 164, 129, 188, 0, 0, 0, 115, 99, 114, 105, 112, 116,
		47, 116, 101, 115, 116, 46, 116, 85, 84, 5, 0, 1, 205, 164,
		225, 90, 80, 75, 5, 6, 0, 0, 0, 0, 3, 0, 3, 0, 246, 0, 0,
		0, 5, 1, 0, 0, 0, 0,
	})
}

func addResource(manager *parcello.Manager) {
	manager.Add([]byte{
		80, 75, 3, 4, 20, 0, 8, 0, 8, 0, 110, 80, 154, 76, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 34, 0, 9, 0, 109, 105, 103, 114,
		97, 116, 105, 111, 110, 47, 48, 48, 48, 54, 48, 53, 50, 52,
		48, 48, 48, 48, 48, 48, 95, 115, 101, 116, 117, 112, 46, 115,
		113, 108, 85, 84, 5, 0, 1, 241, 163, 225, 90, 108, 143, 65,
		75, 195, 64, 16, 133, 239, 251, 43, 222, 81, 161, 1, 83, 47,
		109, 122, 90, 235, 10, 193, 164, 13, 201, 8, 237, 73, 214,
		100, 72, 2, 118, 55, 108, 38, 248, 247, 101, 17, 170, 96,
		231, 52, 60, 190, 7, 223, 75, 18, 232, 69, 124, 210, 179,
		227, 96, 133, 59, 88, 1, 13, 11, 244, 20, 144, 110, 145, 110,
		179, 199, 117, 150, 110, 176, 55, 13, 97, 253, 144, 110, 84,
		146, 160, 250, 100, 59, 51, 58, 15, 231, 5, 237, 96, 93, 207,
		144, 129, 225, 236, 133, 97, 69, 194, 248, 177, 8, 207, 42,
		194, 49, 203, 176, 76, 106, 95, 27, 77, 6, 164, 159, 10, 131,
		252, 5, 135, 35, 193, 156, 242, 134, 26, 92, 198, 62, 88,
		25, 189, 155, 113, 167, 48, 118, 184, 30, 153, 19, 253, 124,
		17, 63, 188, 21, 5, 170, 58, 47, 117, 125, 198, 171, 57, 175,
		20, 58, 158, 219, 48, 78, 177, 124, 3, 94, 41, 180, 129, 227,
		176, 119, 43, 0, 229, 165, 105, 72, 151, 213, 21, 80, 247,
		187, 63, 150, 157, 255, 114, 234, 185, 62, 86, 191, 150, 255,
		12, 119, 234, 59, 0, 0, 255, 255, 80, 75, 7, 8, 149, 201,
		196, 106, 212, 0, 0, 0, 53, 1, 0, 0, 80, 75, 3, 4, 20, 0,
		8, 0, 8, 0, 110, 80, 154, 76, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 34, 0, 9, 0, 109, 105, 103, 114, 97, 116, 105, 111,
		110, 47, 50, 48, 49, 56, 48, 52, 48, 54, 49, 57, 48, 48, 49,
		53, 95, 117, 115, 101, 114, 115, 46, 115, 113, 108, 85, 84,
		5, 0, 1, 241, 163, 225, 90, 84, 205, 193, 74, 195, 64, 16,
		198, 241, 251, 60, 197, 119, 84, 112, 33, 17, 20, 77, 79,
		107, 221, 66, 48, 166, 33, 25, 161, 61, 201, 106, 198, 54,
		80, 55, 101, 119, 22, 95, 95, 170, 30, 234, 117, 254, 195,
		247, 51, 6, 54, 235, 108, 118, 18, 36, 122, 149, 17, 94, 177,
		138, 19, 236, 49, 2, 183, 40, 239, 171, 162, 168, 202, 27,
		44, 221, 192, 184, 46, 202, 59, 50, 6, 221, 65, 124, 18, 140,
		51, 194, 172, 120, 223, 251, 176, 19, 232, 94, 16, 252, 167,
		192, 171, 198, 233, 45, 171, 36, 58, 61, 159, 110, 21, 242,
		145, 150, 189, 179, 236, 192, 246, 161, 113, 200, 73, 98,
		194, 5, 1, 211, 136, 186, 101, 116, 125, 253, 108, 251, 45,
		158, 220, 22, 237, 154, 209, 190, 52, 205, 21, 1, 31, 83,
		76, 250, 250, 51, 204, 110, 195, 255, 218, 193, 159, 39, 186,
		92, 156, 129, 227, 252, 21, 232, 177, 95, 119, 127, 96, 189,
		130, 219, 212, 3, 15, 191, 244, 130, 190, 3, 0, 0, 255, 255,
		80, 75, 7, 8, 161, 196, 255, 151, 191, 0, 0, 0, 251, 0, 0,
		0, 80, 75, 3, 4, 20, 0, 8, 0, 8, 0, 86, 80, 154, 76, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 18, 0, 9, 0, 115, 99, 114, 105,
		112, 116, 47, 99, 111, 109, 109, 97, 110, 100, 46, 115, 113,
		108, 85, 84, 5, 0, 1, 196, 163, 225, 90, 116, 143, 209, 75,
		195, 48, 24, 196, 223, 243, 87, 220, 227, 38, 13, 216, 162,
		34, 21, 25, 213, 125, 226, 160, 110, 210, 102, 238, 81, 130,
		249, 148, 64, 172, 35, 73, 255, 127, 137, 219, 32, 224, 10,
		121, 200, 119, 220, 253, 184, 147, 18, 205, 24, 127, 228,
		23, 15, 236, 117, 100, 3, 29, 177, 99, 131, 102, 239, 81,
		93, 163, 44, 235, 234, 166, 174, 174, 240, 208, 43, 84, 151,
		229, 173, 16, 82, 98, 208, 223, 92, 35, 176, 227, 143, 40,
		181, 115, 114, 12, 236, 131, 232, 169, 165, 71, 133, 11, 60,
		117, 155, 23, 252, 105, 119, 255, 253, 73, 63, 103, 21, 187,
		103, 234, 8, 214, 224, 30, 139, 60, 103, 135, 192, 254, 152,
		91, 173, 123, 234, 20, 86, 107, 181, 57, 164, 48, 179, 166,
		192, 167, 245, 33, 190, 39, 123, 1, 167, 143, 223, 185, 120,
		107, 218, 45, 245, 152, 45, 10, 164, 55, 207, 169, 227, 222,
		232, 200, 7, 234, 246, 117, 217, 40, 194, 105, 133, 202, 120,
		169, 76, 198, 76, 231, 100, 81, 195, 142, 79, 200, 37, 181,
		164, 104, 122, 222, 111, 0, 0, 0, 255, 255, 80, 75, 7, 8,
		185, 201, 214, 166, 208, 0, 0, 0, 121, 1, 0, 0, 80, 75, 3,
		4, 20, 0, 8, 0, 8, 0, 53, 80, 154, 76, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 13, 0, 9, 0, 115, 99, 114, 105, 112, 116, 47,
		116, 101, 120, 116, 46, 116, 85, 84, 5, 0, 1, 134, 163, 225,
		90, 74, 68, 0, 46, 64, 0, 0, 0, 255, 255, 80, 75, 7, 8, 166,
		178, 131, 44, 10, 0, 0, 0, 12, 0, 0, 0, 80, 75, 1, 2, 20,
		3, 20, 0, 8, 0, 8, 0, 110, 80, 154, 76, 149, 201, 196, 106,
		212, 0, 0, 0, 53, 1, 0, 0, 34, 0, 9, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 164, 129, 0, 0, 0, 0, 109, 105, 103, 114, 97, 116, 105,
		111, 110, 47, 48, 48, 48, 54, 48, 53, 50, 52, 48, 48, 48,
		48, 48, 48, 95, 115, 101, 116, 117, 112, 46, 115, 113, 108,
		85, 84, 5, 0, 1, 241, 163, 225, 90, 80, 75, 1, 2, 20, 3, 20,
		0, 8, 0, 8, 0, 110, 80, 154, 76, 161, 196, 255, 151, 191,
		0, 0, 0, 251, 0, 0, 0, 34, 0, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		164, 129, 45, 1, 0, 0, 109, 105, 103, 114, 97, 116, 105, 111,
		110, 47, 50, 48, 49, 56, 48, 52, 48, 54, 49, 57, 48, 48, 49,
		53, 95, 117, 115, 101, 114, 115, 46, 115, 113, 108, 85, 84,
		5, 0, 1, 241, 163, 225, 90, 80, 75, 1, 2, 20, 3, 20, 0, 8,
		0, 8, 0, 86, 80, 154, 76, 185, 201, 214, 166, 208, 0, 0, 0,
		121, 1, 0, 0, 18, 0, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 164, 129,
		69, 2, 0, 0, 115, 99, 114, 105, 112, 116, 47, 99, 111, 109,
		109, 97, 110, 100, 46, 115, 113, 108, 85, 84, 5, 0, 1, 196,
		163, 225, 90, 80, 75, 1, 2, 20, 3, 20, 0, 8, 0, 8, 0, 53,
		80, 154, 76, 166, 178, 131, 44, 10, 0, 0, 0, 12, 0, 0, 0,
		13, 0, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 164, 129, 94, 3, 0, 0,
		115, 99, 114, 105, 112, 116, 47, 116, 101, 120, 116, 46, 116,
		85, 84, 5, 0, 1, 134, 163, 225, 90, 80, 75, 5, 6, 0, 0, 0,
		0, 4, 0, 4, 0, 63, 1, 0, 0, 172, 3, 0, 0, 0, 0,
	})
}
