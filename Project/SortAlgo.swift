//
//  SortAlgo.swift
//  Test
//
//  Created by Eirik Sandberg on 04.05.2017.
//  Copyright © 2017 Eirik Sandberg. All rights reserved.
//

import Foundation
import Test

func generateStringArray(n: Int) -> [String] {
    var array = [String]()
    for index in 0...n {
        var string = randomString(length: 10)
        array.append(string)
    }
    return array
}

//Function to generate random string
// This function is taken from http://stackoverflow.com/questions/26845307/generate-random-alphanumeric-string-in-swift
func randomString(length: Int) -> String {
    
    let letters : NSString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    let len = UInt32(letters.length)
    
    var randomString = ""
    
    for _ in 0 ..< length {
        let rand = arc4random_uniform(len)
        var nextChar = letters.character(at: Int(rand))
        randomString += NSString(characters: &nextChar, length: 1) as String
    }
    
    return randomString
}

func convertArrayToUnicodeValues (array: [String]) -> [Int] {
    var intArray = [Int]()
    for string in array{
        var sum = 0;
        for letter in string.characters{
            sum += letter.hashValue
        }
        intArray.append(sum)
    }
    return intArray
}

// Quicksort algorithm taken from https://github.com/raywenderlich/swift-algorithm-club/blob/master/Quicksort/Quicksort.swift
// MARK: - Hoare partitioning
/*
 Hoare's partitioning scheme.
 The return value is NOT necessarily the index of the pivot element in the
 new array. Instead, the array is partitioned into [low...p] and [p+1...high],
 where p is the return value. The pivot value is placed somewhere inside one
 of the two partitions, but the algorithm doesn't tell you which one or where.
 If the pivot value occurs more than once, then some instances may appear in
 the left partition and others may appear in the right partition.
 Hoare scheme is more efficient than Lomuto's partition scheme; it performs
 fewer swaps.
 */
func partitionHoare<T: Comparable>(_ a: inout [T], low: Int, high: Int) -> Int {
    let pivot = a[low]
    var i = low - 1
    var j = high + 1
    
    while true {
        repeat { j -= 1 } while a[j] > pivot
        repeat { i += 1 } while a[i] < pivot
        
        if i < j {
            swap(&a[i], &a[j])
        } else {
            return j
        }
    }
}

/*
 Recursive, in-place version that uses Hoare's partioning scheme. Because of
 the choice of pivot, this performs badly if the array is already sorted.
 */
func quicksortHoare<T: Comparable>(_ a: inout [T], low: Int, high: Int) {
    if low < high {
        let p = partitionHoare(&a, low: low, high: high)
        quicksortHoare(&a, low: low, high: p)
        quicksortHoare(&a, low: p + 1, high: high)
    }
}
